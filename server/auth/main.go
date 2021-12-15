package main

import (
	"log"
	"net"

	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/wechat"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create zap.logger: %v", err)
	}

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIdResolver: &wechat.Service{},
		Logger:         logger,
	})
	err = s.Serve(lis)
	logger.Fatal("cannot server", zap.Error((err)))
}
