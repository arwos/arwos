package runner

import (
	"context"

	"go.osspkg.com/ioutils/fs"
)

type Runner interface {
}

type _runner struct {
	conf Config
}

func New(c Config) Runner {
	return &_runner{conf: c}
}

func (v *_runner) Up(ctx context.Context) error {
	return nil
}

func (v *_runner) Down() error {
	return nil
}

func (v *_runner) Scan() ([]PluginInfo, error) {
	list, err := fs.SearchFilesByExt(v.conf.Folder, pluginExt)
	if err != nil {
		return nil, err
	}

	for _, item := range list {

	}
}
