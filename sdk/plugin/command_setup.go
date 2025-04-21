/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package plugin

import (
	"context"
	"encoding/json"
	"os"

	"go.osspkg.com/console"
	"go.osspkg.com/events"
	"go.osspkg.com/logx"

	"go.arwos.org/arwos/sdk/rpc"
)

type (
	service interface {
		Up(ctx context.Context) error
		Down() error
	}
)

func (a *_plugin) commandSetup() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("setup", "show addon info")
		setter.Flag(func(flags console.FlagsSetter) {
			flags.String("sock", "unix socket path")
		})
		setter.ExecFunc(func(_ []string, address string) {
			ctx, ctxCancel := context.WithCancel(context.Background())

			logx.SetOutput(os.Stdout)
			logx.SetLevel(logx.LevelDebug)

			go events.OnStopSignal(func() {
				ctxCancel()
			})

			srv := rpc.NewServer(address)

			srv.AddHandler(defaultMethodInfo, func(ctx context.Context, w rpc.Writer, r rpc.Reader) error {
				b, err := json.Marshal(a.info)
				if err != nil {
					return err
				}
				return w.SetField(rpc.FieldBody, b)
			})

			srv.AddHandler(defaultMethodEnvs, func(ctx context.Context, w rpc.Writer, r rpc.Reader) error {
				b, err := json.Marshal(a.envs)
				if err != nil {
					return err
				}
				return w.SetField(rpc.FieldBody, b)
			})

			svc, ok := srv.(service)
			if !ok {
				console.Fatalf("fail start setup command")
			}

			defer func() {
				console.Infof("turn off setup command: %s", address)
				console.FatalIfErr(svc.Down(), "turn off setup command")
			}()

			console.Infof("turn on setup command: %s", address)
			console.FatalIfErr(svc.Up(ctx), "turn on setup command")

			<-ctx.Done()
		})
	})
}
