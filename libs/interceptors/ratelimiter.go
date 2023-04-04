package interceptors

import (
	"context"
	"fmt"
	"route256/libs/logger"
	"time"

	"google.golang.org/grpc"
)

type RateLimiter interface {
	WaitN(ctx context.Context, n int) (err error)
}

type clientRateLimiterInterceptor struct {
	RateLimiter
}

func NewClientRateLimiterInterceptor(rateLimiter RateLimiter) *clientRateLimiterInterceptor {
	return &clientRateLimiterInterceptor{rateLimiter}
}

func (i *clientRateLimiterInterceptor) Intercept(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	i.WaitN(ctx, 1)
	end := time.Now()

	err := invoker(ctx, method, req, reply, cc, opts...)
	logger.Debug(fmt.Sprintf("Invoked RPC method=%s; Wait=%s; Error=%v", method, end.Sub(start), err))

	return err
}
