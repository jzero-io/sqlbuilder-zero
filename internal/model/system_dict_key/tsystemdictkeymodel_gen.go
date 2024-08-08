// Code generated by goctl. DO NOT EDIT.

package system_dict_key

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTSystemDictKeyModel) FindOne(ctx context.Context, id int64) (*TSystemDictKey, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tSystemDictKeyRows, m.table)
	var resp TSystemDictKey
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
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
	query := fmt.Sprintf("select %s from %s where `category_code` = ? limit 1", tSystemDictKeyRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, categoryCode)
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
	query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", tSystemDictKeyRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, uuid)
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
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, tSystemDictKeyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uuid, data.CategoryCode, data.CategoryDesc, data.Sort)
	return ret, err
}

func (m *defaultTSystemDictKeyModel) Update(ctx context.Context, newData *TSystemDictKey) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tSystemDictKeyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Uuid, newData.CategoryCode, newData.CategoryDesc, newData.Sort, newData.Id)
	return err
}

func (m *defaultTSystemDictKeyModel) tableName() string {
	return m.table
}