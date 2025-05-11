package runner

import (
	"context"

	"go.arwos.org/arwos/sdk/manifest"
	"go.osspkg.com/ioutils/fs"
)

type Runner interface {
}

type _runner struct {
	conf RunnerConfig
}

func New(c *ConfigGroup) Runner {
	return &_runner{conf: c.Runner}
}

func (v *_runner) Up(_ context.Context) error {
	list, err := v.Scan()
	if err != nil {
		return err
	}
	return nil
}

func (v *_runner) Down() error {
	return nil
}

func (v *_runner) Scan() ([]PluginInfo, error) {
	list, err := fs.SearchFiles(v.conf.Folder, manifest.FileName)
	if err != nil {
		return nil, err
	}

	result := make([]PluginInfo, 0, len(list))

	for _, item := range list {
		result = append(result, ReadPluginInfo(item))
	}

	return result, nil
}
