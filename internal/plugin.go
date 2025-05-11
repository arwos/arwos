package internal

import (
	"go.arwos.org/arwos/internal/ctrl"
	"go.arwos.org/arwos/internal/pkg/runner"
	"go.arwos.org/arwos/internal/pkg/webcache"
	"go.osspkg.com/goppy/v2/plugins"
)

var Plugins = plugins.Inject(
	plugins.Plugin{
		Config: &ctrl.ConfigGroup{},
		Inject: ctrl.NewAPI,
	},
	plugins.Plugin{
		Config: &runner.ConfigGroup{},
		Inject: runner.New,
	},
	plugins.Plugin{
		Config: &webcache.ConfigGroup{},
		Inject: webcache.NewCache,
	},
)
