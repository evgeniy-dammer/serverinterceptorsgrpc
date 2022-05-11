package main

import (
	"fmt"
	"net"

	"github.com/evgeniy-dammer/serverinterceptorsgrpc/productservice/handlers"
	"github.com/evgeniy-dammer/serverinterceptorsgrpc/productservice/inceptors"
	productservice "github.com/evgeniy-dammer/serverinterceptorsgrpc/productservice/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":1111")

	if err != nil {
		fmt.Println(err)
	}

	defer listen.Close()

	productServ := handlers.ProductServiceServer{}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				inceptors.DateLogInceptor,
				inceptors.MethodLogInceptor,
			),
		),
	)

	productservice.RegisterProductServiceServer(grpcServer, &productServ)

	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
