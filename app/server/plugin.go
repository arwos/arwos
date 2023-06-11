package server

import (
	"github.com/arwos/arwos/pkg/adapters"
	"github.com/arwos/arwos/pkg/agentsrv"
	"github.com/deweppro/goppy/plugins"
	"github.com/deweppro/goppy/plugins/web"
)

var Plugins = plugins.Plugins{}.
	Inject(
		NewApi,
	).
	Inject(
		adapterMainRouter,
		adapters.NewDatabase,
		agentsrv.NewAgentSrv,
	)

func adapterMainRouter(r web.RouterPool) web.Router {
	return r.Main()
}
