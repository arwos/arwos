package webcache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"go.osspkg.com/goppy/v2/web"
	"go.osspkg.com/ioutils/shell"
	"go.osspkg.com/logx"
	"go.osspkg.com/static"
)

type Cache struct {
	conf StaticConfig
	data *static.Cache
	mux  sync.RWMutex
}

func NewCache(c *ConfigGroup) *Cache {
	return &Cache{
		conf: c.Static,
		data: static.New(),
	}
}

func (v *Cache) Up(ctx context.Context) error {
	//if !fs.FileExist(v.conf.Static.Build + "/index.html") {
	if err := v.Rebuild(ctx); err != nil {
		return err
	}
	//}
	return v.Scan()
}

func (v *Cache) Down() error {
	return nil
}

func (v *Cache) Handler(ctx web.Context) {
	v.mux.RLock()
	defer v.mux.RUnlock()

	filename := ctx.Request().RequestURI

	if ext := filepath.Ext(filename); len(ext) == 0 {
		filename = "/index.html"
	}

	if filename != "/index.html" {
		ctx.Header().Set("Cache-Control", "max-age=86400, public")
	}

	if err := v.data.ResponseWrite(ctx.Response(), filename); err != nil {
		logx.Warn("Send static file", "module", "webcache", "file", filename)
	}
}

func (v *Cache) Scan() error {
	s := static.New()
	if err := s.FromDir(v.conf.Build); err != nil {
		return fmt.Errorf("scan webui: %w", err)
	}

	v.mux.Lock()
	defer v.mux.Unlock()

	v.data = s

	return nil
}

func (v *Cache) Rebuild(ctx context.Context) error {
	sh := shell.New()
	sh.SetDir(v.conf.Source)
	err := sh.CallContext(ctx, os.Stdout, "yarn run build")
	if err != nil {
		return fmt.Errorf("failed build static: %w", err)
	}
	return nil
}
