/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package plugin

import (
	"fmt"

	"go.osspkg.com/console"
	"go.osspkg.com/ioutils/data"

	"go.arwos.org/arwos/sdk/env"
	"go.arwos.org/arwos/sdk/manifest"
)

type Plugin interface {
	WithDependencies(deps ...Dep)
	WithDatabase(tags []string)
	Run()
}

type _plugin struct {
	info manifest.Manifest
	envs env.Envs
	deps []any
	conf *data.Buffer
	cli  *console.Console
}

func New(m manifest.Manifest) Plugin {
	a := &_plugin{
		cli:  console.New("arwos.plugin", fmt.Sprintf("author: %s, pkg: %s, des: %s", m.Author, m.Package, m.Description)),
		conf: data.NewBuffer(1024),
		info: m,
	}

	configBase(a.conf)

	return a
}

func (a *_plugin) Run() {
	a.cli.RootCommand(a.commandApp())
	a.cli.AddCommand(a.commandSetup())
	a.cli.Exec()
}
