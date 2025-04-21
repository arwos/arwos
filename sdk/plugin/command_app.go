/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package plugin

import (
	envconf "go.osspkg.com/config/env"
	"go.osspkg.com/console"
	goppy "go.osspkg.com/goppy/v2"
	"go.osspkg.com/goppy/v2/plugins"

	"go.arwos.org/arwos/sdk/rpc"
)

func (a *_plugin) commandApp() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Flag(func(flags console.FlagsSetter) {
			flags.String("sock", "unix socket path")
		})
		setter.ExecFunc(func(_ []string, address string) {
			/*
				Add default deps
			*/
			a.WithDependencies(Dep{Inject: []any{
				rpc.WithRPCServer(address),
			}})

			/*
				Init app
			*/
			app := goppy.New(a.info.Name, a.info.Version, a.info.Description)
			app.ConfigResolvers(envconf.New())
			app.ConfigData(a.conf.String(), ".yaml")

			/*
				Inject app
			*/
			for _, dep := range a.deps {
				switch src := dep.(type) {
				case plugins.Plugin:
					app.Plugins(src)
				case plugins.Plugins:
					app.Plugins(src...)
				default:
					app.Plugins(plugins.Plugin{Inject: src})
				}
			}

			app.Run()
		})
	})
}
