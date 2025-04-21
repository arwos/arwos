package runner

import (
	"go.arwos.org/arwos/sdk/env"
	"go.arwos.org/arwos/sdk/manifest"
)

const pluginExt = ".arwos"

type PluginInfo struct {
	Path string
	Hash string
	Meta manifest.Manifest
	Envs env.Envs
}
