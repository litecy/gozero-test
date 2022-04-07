package logic

import (
	"context"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"

	"gozero-test/internal/svc"
	"gozero-test/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmptyCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyCallLogic {
	return &EmptyCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmptyCallLogic) EmptyCall(in *model.Empty) (*model.Empty, error) {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999)

	var idForCaller3 int64

	{
		//caller 1
		tp := rand.Intn(100)
		ret, err := l.svcCtx.ATestModel.Insert(l.ctx, &model.Atest{
			Code: int64(code),
			Tp:   int64(tp),
		})

		if err != nil {
			l.Errorf("insert failed, %+v", err)
			return nil, err
		}

		lastId, _ := ret.LastInsertId()

		idForCaller3 = lastId

		l.Infof("[direct]new added code=%d, id=%d, type=%d", code, lastId, tp)

		firstQuery, err := l.svcCtx.ATestModel.FindOne(l.ctx, lastId)
		if err != nil {
			l.Errorf("query first failed, %+v", err)
			return nil, err
		}
		l.Infof("[query-1]new added code=%d, id=%d, type=%d", firstQuery.Code, firstQuery.Id, firstQuery.Tp)
	}
	{
		//caller 2
		// assume the next call is from another function and only pass the code parameter
		time.After(time.Second * 1)
		tp := rand.Intn(100)
		err := l.svcCtx.ATestModel.UpdateTypeByCode(l.ctx, int64(code), int64(tp))
		if err != nil {
			l.Errorf("UpdateTypeByCode failed, %+v", err)
			return nil, err
		}

		l.Infof("type is updated to %d for code: %d", tp, code)
	}

	{
		//caller 3
		//later, the other caller called FindOne by the id
		time.After(time.Second * 1)
		ret, err := l.svcCtx.ATestModel.FindOne(l.ctx, idForCaller3)
		if err != nil {
			l.Errorf("UpdateTypeByCode failed, %+v", err)
			return nil, err
		}
		l.Infof("[query-3]code=%d, id=%d, type=%d", ret.Code, ret.Id, ret.Tp)
	}

	{
		//caller 4
		// find by code
		time.After(time.Second * 1)
		ret, err := l.svcCtx.ATestModel.FindOneByCode(l.ctx, int64(code))
		if err != nil {
			l.Errorf("FindOneByCode failed, %+v", err)
			return nil, err
		}
		l.Infof("[query-4]code=%d, id=%d, type=%d", ret.Code, ret.Id, ret.Tp)
	}

	return &model.Empty{}, nil
}
