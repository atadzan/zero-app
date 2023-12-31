package main

import (
	"flag"
	"fmt"

	"zero-app/rpc/auth/auth"
	"zero-app/rpc/auth/internal/config"
	"zero-app/rpc/auth/internal/server"
	"zero-app/rpc/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		auth.RegisterAuthServer(grpcServer, server.NewAuthServer(ctx))

		//if c.Mode == service.DevMode || c.Mode == service.TestMode {
		//	reflection.Register(grpcServer)
		//}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
