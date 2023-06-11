package server

import (
	"github.com/arwos/arwos/pkg/adapters"
	"github.com/arwos/arwos/pkg/agentsrv"
	"github.com/deweppro/go-sdk/app"
	"github.com/deweppro/goppy/plugins/web"
)

type Api struct {
	route  web.Router
	db     adapters.DB
	agents *agentsrv.AgentSrv
}

func NewApi(r web.Router, db adapters.DB, as *agentsrv.AgentSrv) *Api {
	return &Api{
		route:  r,
		db:     db,
		agents: as,
	}
}

func (v *Api) Up(_ app.Context) error {
	v.agents.InjectApi(v.route)
	return nil
}

func (v *Api) Down() error {
	return nil
}
