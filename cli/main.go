package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gozero-test/model"
)

func main() {
	cli, err := zrpc.NewClientWithTarget("127.0.0.1:8080")
	if err != nil {
		logx.Errorf("new client failed, %+v", err)
		return
	}
	svc := model.NewEmptySvcClient(cli.Conn())
	_, err = svc.EmptyCall(context.TODO(), &model.Empty{})
	if err != nil {
		logx.Errorf("empty call failed, %+v", err)
		return
	}
}
