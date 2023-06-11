package main

import (
	"github.com/arwos/arwos/app/agent"
	"github.com/deweppro/goppy"
	"github.com/deweppro/goppy/plugins/web"
)

func main() {
	app := goppy.New()
	app.WithConfig("./config.agent.yaml") // Reassigned via the `--config` argument when run via the console.
	app.Plugins(
		web.WithHTTPClient(),
		web.WithWebsocketClient(),
	)
	app.Plugins(agent.Plugins...)
	app.Run()
}
