package system_dict_key

import (
	"context"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TSystemDictKeyModel = (*customTSystemDictKeyModel)(nil)

type (
	// TSystemDictKeyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTSystemDictKeyModel.
	TSystemDictKeyModel interface {
		tSystemDictKeyModel
		WithSession(session sqlx.Session) TSystemDictKeyModel

		BulkInsert(ctx context.Context, datas []*TSystemDictKey) error
		Find(ctx context.Context, conds ...condition.Condition) ([]*TSystemDictKey, error)
		Page(ctx context.Context, conds ...condition.Condition) ([]*TSystemDictKey, int64, error)
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
