package adapters

import (
	"github.com/deweppro/go-sdk/orm"
	"github.com/deweppro/goppy/plugins/database"
)

type (
	adapterDB struct {
		database database.MySQL
	}

	DB interface {
		Pool() orm.Stmt
	}
)

func NewDatabase(db database.MySQL) DB {
	return &adapterDB{
		database: db,
	}
}

func (v *adapterDB) Pool() orm.Stmt {
	return v.database.Pool("main")
}
