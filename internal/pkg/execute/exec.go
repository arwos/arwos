package execute

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"go.arwos.org/arwos/internal/pkg/runner"
	"go.arwos.org/arwos/sdk/env"
	"go.arwos.org/arwos/sdk/manifest"
	"go.arwos.org/arwos/sdk/rpc"
	"go.osspkg.com/ioutils/shell"
)

type Execute struct {
	sock string
	svc  bool
	log  io.Writer
	cli  rpc.Client
	info runner.PluginInfo
	envs []runner.PluginEnv
}

func New(i runner.PluginInfo, e []runner.PluginEnv, l io.Writer) (*Execute, error) {
	ex := &Execute{
		sock: fmt.Sprintf("/tmp/%s.sock", uuid.New()),
		info: i,
		envs: e,
		log:  l,
		svc:  i.Meta.Type == manifest.TypeService,
	}

	if ex.IsService() {
		var err error
		ex.cli, err = rpc.NewClient(ex.sock, rpc.ClientMaxConns(10))
		if err != nil {
			return nil, err
		}
	}

	return ex, nil
}

func (e *Execute) IsService() bool {
	return e.svc
}

func (e *Execute) Watch(ctx context.Context) {
	if !e.IsService() {
		<-ctx.Done()
	}

	sh := shell.New()
	sh.SetDir(e.info.Root)
	for _, item := range e.envs {
		sh.SetEnv(item.Key, item.Value)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := sh.CallContext(ctx, e.log, fmt.Sprintf("%s --sock=%s", e.info.Meta.Path, e.sock)); err != nil {
				io.WriteString(e.log, err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}
}

const CallEnvPrefix = "INPUT_"

func (e *Execute) Call(ctx context.Context, req rpc.Request) rpc.Response {
	if !e.IsService() {
		sh := shell.New()
		sh.SetDir(e.info.Root)
		sh.SetShell("/bin/sh", "e", "c")
		for _, item := range e.envs {
			sh.SetEnv(item.Key, item.Value)
		}

		if !req.Deadline().IsZero() {
			var ctxCancel context.CancelFunc
			ctx, ctxCancel = context.WithDeadline(ctx, req.Deadline())
			defer func() { ctxCancel() }()
		}

		tr, ok := asTransport(req)
		if !ok {
			tr.SetCode(500)
			tr.SetError(fmt.Errorf("invalid request type"))
			return tr
		}

		sh.SetEnv(CallEnvPrefix+"METHOD", tr.Method())
		sh.SetEnv(CallEnvPrefix+"BODY", tr.String())
		for _, item := range tr.GetCtxKeys() {
			if val, ok := tr.GetCtx(item); ok {
				sh.SetEnv(CallEnvPrefix+"CTX_"+env.CanonicalName(item), val)
			}
		}
		tr.SoftReset()

		b, err := sh.Call(ctx, e.info.Meta.Path)
		if err != nil {
			tr.SetCode(400)
			tr.SetError(err)
		} else {
			tr.Write(b) //nolint:errcheck
		}

		tr.Seek(0, 0) //nolint:errcheck

		return tr
	}

	if e.cli == nil {
		tr, _ := asTransport(req)
		tr.SoftReset()
		tr.SetCode(500)
		tr.SetError(fmt.Errorf("client not initealized"))
		return tr
	}

	return e.cli.Call(ctx, req)
}

func asTransport(req rpc.Request) (rpc.Transport, bool) {
	tr, ok := req.(rpc.Transport)
	if !ok {
		return rpc.NewTransport(), false
	}
	return tr, true
}
