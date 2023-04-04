package metrics

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	timeStart := time.Now()
	res, err := handler(ctx, req)
	elapsed := time.Since(timeStart)

	label := "Success"
	if err != nil {
		label = "Error"
	}
	RequestsCounter.Inc()
	ResponseCounter.WithLabelValues(label).Inc()
	HistogramResponseTime.WithLabelValues(label).Observe(elapsed.Seconds())

	if err != nil {
		return nil, err
	}

	return res, nil
}
