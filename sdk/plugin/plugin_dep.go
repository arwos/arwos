/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package plugin

import "go.arwos.org/arwos/sdk/env"

type Dep struct {
	Inject []any
	Config string
	Envs   []env.Env
}

func (a *_plugin) WithDependencies(deps ...Dep) {
	for _, dep := range deps {
		a.deps = append(a.deps, dep.Inject...)
		a.envs = append(a.envs, dep.Envs...)
		confWrite(a.conf, dep.Config)
	}
}
