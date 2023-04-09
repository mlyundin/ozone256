package interceptors

import (
	"context"
	"fmt"
	"time"

	"route256/libs/logger"

	"google.golang.org/grpc"
)

const dateLayout = "2006-01-02"

// LoggingInterceptor ...
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	logger.Info(fmt.Sprintf("[gRPC] %s: %s --- %v", time.Now().Format(dateLayout), info.FullMethod, req))

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(fmt.Sprintf("[gRPC] %s: %s --- %v\n", time.Now().Format(dateLayout), info.FullMethod, err))
		return nil, err
	}

	logger.Info(fmt.Sprintf("[gRPC] %s: %s --- %v\n", time.Now().Format(dateLayout), info.FullMethod, res))

	return res, nil
}
