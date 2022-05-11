package inceptors

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func DateLogInceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	today := time.Now()

	fmt.Println("Access product service date: " + today.Format("01/02/2006 15:04:05"))
	return handler(ctx, req)
}