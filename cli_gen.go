package main

//go:generate goctl rpc protoc atest.proto --go_out=./ --go-grpc_out=./ --zrpc_out=./

// go:generate goctl model mysql datasource --url=root:123456@(localhost:3306)/mytest -t atest -dir=./model --style go_zero
//go:generate goctl model mysql ddl --src ./atest.sql -dir=./model --style go_zero -c
