/*
 *  Copyright (c) 2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"context"
	"net/http"

	"go.osspkg.com/goppy/v2/orm"

	"go.arwos.org/arwos/sdk/manifest"
	"go.arwos.org/arwos/sdk/plugin"
	"go.arwos.org/arwos/sdk/rpc"
)

func main() {

	app := plugin.New(manifest.Model{
		Name:        "simple",
		Package:     "com.example.simple",
		Description: "Simple Addon",
		Author:      "UserName <user@example.com>",
		Version:     "v0.0.0-dev",
		Type:        manifest.TypeService,
		Links: []manifest.Link{
			{Url: "https://example.com", Description: "Web"},
			{Url: "https://tg/example", Description: "Telegram"},
		},
		Menu: manifest.Menu{
			Group: "Application",
			Title: "Simple App",
		},
	})

	app.WithDatabase([]string{"master", "slave"})

	app.WithDependencies(plugin.Dep{Inject: []any{NewController}})
	app.Run()
}

type Controller struct {
	orm orm.ORM
}

func NewController(db orm.ORM, srv rpc.Server) error {
	c := &Controller{orm: db}
	srv.AddHandler("com.example.simple.hello", c.Hello)
	return nil
}

func (c *Controller) Hello(_ context.Context, w rpc.Writer, r rpc.Reader) error {
	w.SetCode(http.StatusOK)
	return w.SetField(rpc.FieldBody, []byte("Hello"))
}
