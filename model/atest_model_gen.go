// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	atestFieldNames          = builder.RawFieldNames(&Atest{})
	atestRows                = strings.Join(atestFieldNames, ",")
	atestRowsExpectAutoSet   = strings.Join(stringx.Remove(atestFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	atestRowsWithPlaceHolder = strings.Join(stringx.Remove(atestFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAtestIdPrefix   = "cache:atest:id:"
	cacheAtestCodePrefix = "cache:atest:code:"
)

type (
	atestModel interface {
		Insert(ctx context.Context, data *Atest) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Atest, error)
		FindOneByCode(ctx context.Context, code int64) (*Atest, error)
		Update(ctx context.Context, data *Atest) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAtestModel struct {
		sqlc.CachedConn
		table string
	}

	Atest struct {
		Id   int64 `db:"id"`
		Code int64 `db:"code"` // unique code
		Tp   int64 `db:"type"` // type
	}
)

func newAtestModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAtestModel {
	return &defaultAtestModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`atest`",
	}
}

func (m *defaultAtestModel) Insert(ctx context.Context, data *Atest) (sql.Result, error) {
	atestIdKey := fmt.Sprintf("%s%v", cacheAtestIdPrefix, data.Id)
	atestCodeKey := fmt.Sprintf("%s%v", cacheAtestCodePrefix, data.Code)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, atestRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Code, data.Tp)
	}, atestIdKey, atestCodeKey)
	return ret, err
}

func (m *defaultAtestModel) FindOne(ctx context.Context, id int64) (*Atest, error) {
	atestIdKey := fmt.Sprintf("%s%v", cacheAtestIdPrefix, id)
	var resp Atest
	err := m.QueryRowCtx(ctx, &resp, atestIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", atestRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAtestModel) FindOneByCode(ctx context.Context, code int64) (*Atest, error) {
	atestCodeKey := fmt.Sprintf("%s%v", cacheAtestCodePrefix, code)
	var resp Atest
	err := m.QueryRowIndexCtx(ctx, &resp, atestCodeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `code` = ? limit 1", atestRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, code); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAtestModel) Update(ctx context.Context, data *Atest) error {
	atestIdKey := fmt.Sprintf("%s%v", cacheAtestIdPrefix, data.Id)
	atestCodeKey := fmt.Sprintf("%s%v", cacheAtestCodePrefix, data.Code)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, atestRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Code, data.Tp, data.Id)
	}, atestIdKey, atestCodeKey)
	return err
}

func (m *defaultAtestModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	atestIdKey := fmt.Sprintf("%s%v", cacheAtestIdPrefix, id)
	atestCodeKey := fmt.Sprintf("%s%v", cacheAtestCodePrefix, data.Code)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, atestIdKey, atestCodeKey)
	return err
}

func (m *defaultAtestModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAtestIdPrefix, primary)
}

func (m *defaultAtestModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", atestRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAtestModel) tableName() string {
	return m.table
}
