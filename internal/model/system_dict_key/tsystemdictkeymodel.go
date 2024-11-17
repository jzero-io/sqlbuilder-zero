package system_dict_key

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSystemDictKeyModel = (*customTSystemDictKeyModel)(nil)

type (
	// TSystemDictKeyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSystemDictKeyModel.
	TSystemDictKeyModel interface {
		tSystemDictKeyModel
	}

	customTSystemDictKeyModel struct {
		*defaultTSystemDictKeyModel
	}
)

// NewTSystemDictKeyModel returns a model for the database table.
func NewTSystemDictKeyModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) TSystemDictKeyModel {
	return &customTSystemDictKeyModel{
		defaultTSystemDictKeyModel: newTSystemDictKeyModel(conn, op...),
	}
}
