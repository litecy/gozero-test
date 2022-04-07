// Code generated by goctl. DO NOT EDIT!
// Source: atest.proto

package server

import (
	"context"

	"gozero-test/internal/logic"
	"gozero-test/internal/svc"
	"gozero-test/model"
)

type EmptySvcServer struct {
	svcCtx *svc.ServiceContext
	model.UnimplementedEmptySvcServer
}

func NewEmptySvcServer(svcCtx *svc.ServiceContext) *EmptySvcServer {
	return &EmptySvcServer{
		svcCtx: svcCtx,
	}
}

func (s *EmptySvcServer) EmptyCall(ctx context.Context, in *model.Empty) (*model.Empty, error) {
	l := logic.NewEmptyCallLogic(ctx, s.svcCtx)
	return l.EmptyCall(in)
}