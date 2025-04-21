package settings

import (
	"time"

	"go.arwos.org/arwos/sdk/manifest"
)

//go:generate goppy gen --type=orm:pgsql --db-read=slave --db-write=master --index=1000 --sql-dir=../../../db-migrations

//gen:orm table=plugins index=uniq:pkg,ver index=uniq:hash
type PluginModel struct {
	Id      int64  // col=id index=pk
	Package string // col=pkg len=300
	Version string // col=ver len=50
	Hash    string // col=hash len=64

	Active bool               // col=active
	Meta   *manifest.Manifest // col=meta

	CreatedAt time.Time // col=created_at auto=c:time.Now()
	UpdatedAt time.Time // col=updated_at auto=time.Now()
}

//gen:orm table=envs index=uniq:plugin_id,key
type EnvModel struct {
	Id       int64 // col=id index=pk
	PluginId int64 // col=plugin_id index=fk:plugins.id

	Key     string // col=key
	Value   string // col=value
	Default string // col=default_value
	Desc    string // col=desc

	CreatedAt time.Time // col=created_at auto=c:time.Now()
	UpdatedAt time.Time // col=updated_at auto=time.Now()
}
