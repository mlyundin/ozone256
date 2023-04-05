package tracing

import (
	"route256/libs/logger"

	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

func Init(serviceName string) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}
	envcfg, err := cfg.FromEnv()
	if err != nil {
		logger.Fatal("Cannot read env var", zap.Error(err))
	}

	_, err = envcfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("Cannot init tracing", zap.Error(err))
	}
}
