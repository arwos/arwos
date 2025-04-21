/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package rpc

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"go.osspkg.com/ioutils/cache"
	"go.osspkg.com/ioutils/fs"
	"go.osspkg.com/logx"
	"go.osspkg.com/network/errs"
	"go.osspkg.com/network/server"
	"go.osspkg.com/syncing"

	errs2 "go.arwos.org/arwos/sdk/rpc/errs"
	"go.arwos.org/arwos/sdk/rpc/internal"
)

type _server struct {
	address  string
	serv     server.Server
	wg       syncing.Group
	handlers cache.TCache[string, Handler]
	guards   cache.TCache[string, Guard]
}

func NewServer(address string) Server {
	return &_server{
		address: address,
		serv: server.New(server.Config{
			Address: address,
			Network: "unix",
		}),
		handlers: cache.NewWithReplace[string, Handler](),
		guards:   cache.NewWithReplace[string, Guard](),
		wg:       syncing.NewGroup(),
	}
}

func (s *_server) AddHandler(method string, call Handler) {
	s.handlers.Set(method, call)
}

func (s *_server) AddGuard(call Guard, methods ...string) {
	for _, method := range methods {
		s.guards.Set(method, call)
	}
}

func (s *_server) Up(ctx context.Context) error {
	if fs.FileExist(s.address) {
		if err := os.Remove(s.address); err != nil {
			return err
		}
	}

	s.serv.HandleFunc(s.handling)

	s.wg.Background(func() {
		if err := s.serv.ListenAndServe(ctx); err != nil && !errs.IsClosed(err) {
			logx.Error("Run Unix Server", "err", err)
		}
	})

	return nil
}

func (s *_server) Down() error {
	s.wg.Wait()

	if fs.FileExist(s.address) {
		if err := os.Remove(s.address); err != nil {
			return err
		}
	}

	return nil
}

func (s *_server) handling(ctx context.Context, w io.Writer, r io.Reader, _ net.Addr) {
	rt, wt := internal.Pool.Get(), internal.Pool.Get()
	defer func() {
		internal.Pool.Put(rt)
		internal.Pool.Put(wt)
	}()

	defer func() {
		if err := wt.Encode(w); err != nil {
			logx.Error("RPC Encode", "method", rt.Method(), "err", err.Error())
		}
	}()

	if err := rt.Decode(r); err != nil {
		logx.Error("RPC Decode", "method", rt.Method(), "err", err.Error())

		wt.SetCode(http.StatusInternalServerError)
		wt.SetError(err)
		return
	}

	wt.SetMethod(rt.Method())

	defer func() {
		if e := recover(); e != nil {
			wt.SoftReset()
			wt.SetCode(http.StatusInternalServerError)
			wt.SetError(fmt.Errorf("RPC Handler Panic: %+v", e))
			return
		}
	}()

	ctx, ctxCancel := context.WithDeadline(ctx, rt.Deadline())
	defer func() { ctxCancel() }()

	if g, ok := s.guards.Get(rt.Method()); ok {
		if err := g(ctx, rt); err != nil {
			logx.Error("RPC Guard", "method", rt.Method(), "err", err.Error())
			wt.SetCode(http.StatusBadRequest)
			wt.SetError(err)
			return
		}
	}

	h, ok := s.handlers.Get(rt.Method())
	if !ok {
		logx.Error("RPC Call", "method", rt.Method(), "err", errs2.ErrMethodNotAllowed.Error())
		wt.SetCode(http.StatusMethodNotAllowed)
		wt.SetError(errs2.ErrMethodNotAllowed)
		return
	}

	if err := h(ctx, wt, rt); err != nil {
		logx.Error("RPC Call", "method", rt.Method(), "err", err.Error())
		wt.SoftReset()
		wt.SetCode(http.StatusInternalServerError)
		wt.SetError(err)
		return
	}
}
