package inceptors

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func MethodLogInceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("Method: " + info.FullMethod)

	return handler(ctx, req)
}
