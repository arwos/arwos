package runner

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"

	"go.arwos.org/arwos/sdk/env"
	"go.arwos.org/arwos/sdk/manifest"
	"go.osspkg.com/errors"
	"go.osspkg.com/ioutils/fs"
	"go.osspkg.com/ioutils/hash"
	"gopkg.in/yaml.v3"
)

func ReadPluginInfo(filename string) PluginInfo {
	dir := filepath.Dir(filename)

	data := PluginInfo{
		Root: dir,
	}

	manifestPath := dir + "/" + manifest.FileName
	b, err := os.ReadFile(manifestPath)
	if err != nil {
		data.Err = errors.Wrap(data.Err, fmt.Errorf("read manifest `%s`: %w", manifestPath, err))
		return data
	}
	if err = yaml.Unmarshal(b, &data.Meta); err != nil {
		data.Err = errors.Wrap(data.Err, fmt.Errorf("decode manifest `%s`: %w", manifestPath, err))
		return data
	}

	envPath := dir + "/" + env.FileName
	if b, err = os.ReadFile(envPath); err == nil {
		if err = yaml.Unmarshal(b, &data.Envs); err != nil {
			data.Err = errors.Wrap(data.Err, fmt.Errorf("decode envs `%s`: %w", envPath, err))
			return data
		}
	}

	binPath := dir + "/" + data.Meta.Path
	if !fs.FileExist(binPath) {
		data.Err = errors.Wrap(data.Err, fmt.Errorf("plugin not found: `%s`", binPath))
		return data
	}

	data.Hash, err = hash.Create(binPath, sha1.New())
	if err != nil {
		data.Err = errors.Wrap(data.Err, fmt.Errorf("calc hash: %w", err))
		return data
	}

	return data
}
