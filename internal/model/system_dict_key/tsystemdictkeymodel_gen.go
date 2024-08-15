// Code generated by goctl. DO NOT EDIT.

package system_dict_key

import (
	"context"
	"database/sql"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tSystemDictKeyFieldNames          = builder.RawFieldNames(&TSystemDictKey{})
	tSystemDictKeyRows                = strings.Join(tSystemDictKeyFieldNames, ",")
	tSystemDictKeyRowsExpectAutoSet   = strings.Join(stringx.Remove(tSystemDictKeyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tSystemDictKeyRowsWithPlaceHolder = strings.Join(stringx.Remove(tSystemDictKeyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tSystemDictKeyModel interface {
		Insert(ctx context.Context, data *TSystemDictKey) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TSystemDictKey, error)
		FindOneByCategoryCode(ctx context.Context, categoryCode string) (*TSystemDictKey, error)
		FindOneByUuid(ctx context.Context, uuid string) (*TSystemDictKey, error)
		Update(ctx context.Context, data *TSystemDictKey) error
		Delete(ctx context.Context, id int64) error

		// custom interface generated by jzero
		BulkInsert(ctx context.Context, datas []*TSystemDictKey) error
		FindByCondition(ctx context.Context, conds ...condition.Condition) ([]*TSystemDictKey, error)
		FindOneByCondition(ctx context.Context, conds ...condition.Condition) (*TSystemDictKey, error)
		PageByCondition(ctx context.Context, conds ...condition.Condition) ([]*TSystemDictKey, int64, error)
	}

	defaultTSystemDictKeyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TSystemDictKey struct {
		Id           int64  `db:"id"`
		Uuid         string `db:"uuid"`          // uuid
		CategoryCode string `db:"category_code"` // 字典类型
		CategoryDesc string `db:"category_desc"` // 字典描述
		Sort         int64  `db:"sort"`          // 排序（升序）
	}
)

func newTSystemDictKeyModel(conn sqlx.SqlConn) *defaultTSystemDictKeyModel {
	return &defaultTSystemDictKeyModel{
		conn:  conn,
		table: "`T_system_dict_key`",
	}
}

func (m *defaultTSystemDictKeyModel) Delete(ctx context.Context, id int64) error {
	sb := sqlbuilder.DeleteFrom(m.table)
	sb.Where(sb.EQ("`id`", id))
	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}

func (m *defaultTSystemDictKeyModel) FindOne(ctx context.Context, id int64) (*TSystemDictKey, error) {
	sb := sqlbuilder.Select(tSystemDictKeyRows).From(m.table)
	sb.Where(sb.EQ("`id`", id))
	sb.Limit(1)
	sql, args := sb.Build()
	var resp TSystemDictKey
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTSystemDictKeyModel) FindOneByCategoryCode(ctx context.Context, categoryCode string) (*TSystemDictKey, error) {
	var resp TSystemDictKey
	sb := sqlbuilder.Select(tSystemDictKeyRows).From(m.table)
	// patch
	sb.Where(sb.EQ(strings.Split(strings.ReplaceAll("`category_code` = ?", " ", ""), "=")[0], categoryCode))
	sb.Limit(1)

	sql, args := sb.Build()
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)

	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTSystemDictKeyModel) FindOneByUuid(ctx context.Context, uuid string) (*TSystemDictKey, error) {
	var resp TSystemDictKey
	sb := sqlbuilder.Select(tSystemDictKeyRows).From(m.table)
	// patch
	sb.Where(sb.EQ(strings.Split(strings.ReplaceAll("`uuid` = ?", " ", ""), "=")[0], uuid))
	sb.Limit(1)

	sql, args := sb.Build()
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)

	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTSystemDictKeyModel) Insert(ctx context.Context, data *TSystemDictKey) (sql.Result, error) {
	sql, args := sqlbuilder.NewInsertBuilder().
		InsertInto(m.table).
		Cols(tSystemDictKeyRowsExpectAutoSet).
		Values(data.Uuid, data.CategoryCode, data.CategoryDesc, data.Sort).Build()
	ret, err := m.conn.ExecCtx(ctx, sql, args...)
	return ret, err
}

func (m *defaultTSystemDictKeyModel) Update(ctx context.Context, newData *TSystemDictKey) error {
	sb := sqlbuilder.Update(m.table)
	split := strings.Split(tSystemDictKeyRowsExpectAutoSet, ",")
	var assigns []string
	for _, s := range split {
		assigns = append(assigns, sb.Assign(s, nil))
	}
	sb.Set(assigns...)
	sb.Where(sb.EQ("`id`", nil))
	sql, _ := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, newData.Uuid, newData.CategoryCode, newData.CategoryDesc, newData.Sort, newData.Id)
	return err
}

func (m *defaultTSystemDictKeyModel) tableName() string {
	return m.table
}

func (m *customTSystemDictKeyModel) BulkInsert(ctx context.Context, datas []*TSystemDictKey) error {
	sb := sqlbuilder.InsertInto(m.table)
	sb.Cols(tSystemDictKeyRowsExpectAutoSet)
	for _, data := range datas {
		sb.Values(data.Uuid, data.CategoryCode, data.CategoryDesc, data.Sort)
	}
	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}

func (m *customTSystemDictKeyModel) FindByCondition(ctx context.Context, conds ...condition.Condition) ([]*TSystemDictKey, error) {
	sb := sqlbuilder.Select(tSystemDictKeyFieldNames...).From(m.table)
	condition.Apply(sb, conds...)
	sql, args := sb.Build()

	var resp []*TSystemDictKey
	err := m.conn.QueryRowsCtx(ctx, &resp, sql, args...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customTSystemDictKeyModel) FindOneByCondition(ctx context.Context, conds ...condition.Condition) (*TSystemDictKey, error) {
	sb := sqlbuilder.Select(tSystemDictKeyFieldNames...).From(m.table)
	condition.Apply(sb, conds...)
	sb.Limit(1)
	sql, args := sb.Build()

	var resp TSystemDictKey
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customTSystemDictKeyModel) PageByCondition(ctx context.Context, conds ...condition.Condition) ([]*TSystemDictKey, int64, error) {
	sb := sqlbuilder.Select(tSystemDictKeyFieldNames...).From(m.table)
	countsb := sqlbuilder.Select("count(*)").From(m.table)

	condition.Apply(sb, conds...)
	condition.Apply(countsb, conds...)

	var resp []*TSystemDictKey

	sql, args := sb.Build()
	err := m.conn.QueryRowsCtx(ctx, &resp, sql, args...)
	if err != nil {
		return nil, 0, err
	}

	// get total
	var total int64
	sql, args = countsb.Build()
	err = m.conn.QueryRowCtx(ctx, &total, sql, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}
