/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"go.arwos.org/arwos/internal"
	"go.osspkg.com/goppy/v2"
	"go.osspkg.com/goppy/v2/auth"
	"go.osspkg.com/goppy/v2/orm"
	"go.osspkg.com/goppy/v2/web"
)

var Version string = "0.0.0-dev"

func main() {
	app := goppy.New("arwos.platform", Version, "Arwos Platform")

	app.Plugins(
		orm.WithMigration(),
		orm.WithORM(),
		orm.WithPgsqlClient(),
	)

	app.Plugins(
		web.WithServer(),
		web.WithClient(),
		auth.WithJWT(),
	)

	app.Plugins(internal.Plugins...)

	app.Run()
}
