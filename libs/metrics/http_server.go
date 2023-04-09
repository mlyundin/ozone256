package metrics

import (
	"fmt"
	"net/http"
	"route256/libs/logger"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunHttpServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		logger.Fatal("Unable to start http server for metrics")
	}
}
