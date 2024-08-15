package system_dict_key

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSystemDictKeyModel = (*customTSystemDictKeyModel)(nil)

type (
	// TSystemDictKeyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSystemDictKeyModel.
	TSystemDictKeyModel interface {
		tSystemDictKeyModel
		WithSession(session sqlx.Session) TSystemDictKeyModel
	}

	customTSystemDictKeyModel struct {
		*defaultTSystemDictKeyModel
	}
)

// NewTSystemDictKeyModel returns a model for the database table.
func NewTSystemDictKeyModel(conn sqlx.SqlConn) TSystemDictKeyModel {
	return &customTSystemDictKeyModel{
		defaultTSystemDictKeyModel: newTSystemDictKeyModel(conn),
	}
}

func (m *customTSystemDictKeyModel) WithSession(session sqlx.Session) TSystemDictKeyModel {
	return NewTSystemDictKeyModel(sqlx.NewSqlConnFromSession(session))
}
