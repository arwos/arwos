package runner

import (
	"go.arwos.org/arwos/sdk/env"
	"go.arwos.org/arwos/sdk/manifest"
)

type PluginInfo struct {
	Root string
	Hash string
	Meta manifest.Model
	Envs env.Model
	Err  error
}

type PluginEnv struct {
	Key   string
	Value string
}
