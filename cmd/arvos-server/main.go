package main

import (
	"github.com/arwos/arwos/app/server"
	"github.com/deweppro/goppy"
	"github.com/deweppro/goppy/plugins/database"
	"github.com/deweppro/goppy/plugins/web"
)

func main() {
	app := goppy.New()
	app.WithConfig("./config.server.yaml") // Reassigned via the `--config` argument when run via the console.
	app.Plugins(
		web.WithHTTP(),
		web.WithHTTPDebug(),
		web.WithWebsocketServerPool(),
		database.WithMySQL(),
	)
	app.Plugins(server.Plugins...)
	app.Run()
}
