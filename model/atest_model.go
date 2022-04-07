package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AtestModel = (*customAtestModel)(nil)

type (
	// AtestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAtestModel.
	AtestModel interface {
		atestModel
		UpdateTypeByCode(ctx context.Context, code int64, typ int64) error
	}

	customAtestModel struct {
		*defaultAtestModel
	}
)

// NewAtestModel returns a model for the database table.
func NewAtestModel(conn sqlx.SqlConn, c cache.CacheConf) AtestModel {
	return &customAtestModel{
		defaultAtestModel: newAtestModel(conn, c),
	}
}

func (m *customAtestModel) UpdateTypeByCode(ctx context.Context, code int64, typ int64) error {
	// FIXME: where to find the id value for special code , query again after mysql exec or set atestIdKey to `cacheAtestIdPrefix*`
	// atestIdKey := fmt.Sprintf("%s%v", cacheAtestIdPrefix, data.Id)
	atestCodeKey := fmt.Sprintf("%s%v", cacheAtestCodePrefix, code)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `type` = ? where `code` = ?", m.table)
		return conn.ExecCtx(ctx, query, typ, code)
	},
		//atestIdKey,
		atestCodeKey)
	return err
}
