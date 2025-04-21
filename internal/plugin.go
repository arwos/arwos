package internal

import (
	"go.arwos.org/arwos/internal/pkg/runner"
	"go.osspkg.com/goppy/v2/plugins"
)

var Plugins = plugins.Inject(
	plugins.Plugin{
		Config: &runner.ConfigGroup{},
		Inject: func(c *runner.ConfigGroup) runner.Runner {
			return runner.New(c.Config)
		},
	},
)
