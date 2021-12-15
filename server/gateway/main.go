package main

import (
	"context"
	"log"
	"net/http"

	authpb "coolcar/auth/api/gen/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	c := context.Background()

	c, cancel := context.WithCancel(c)

	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
			UseProtoNames:  true,
		},
	}))

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(c, mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Fatalf("cannot register auth service: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))
}
