package tracing

import (
	"fmt"
	"log"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func InitTracer() {
	// TRACER INIT
	jaegerHost := os.Getenv("JAEGER_AGENT_HOST")

	if jaegerHost != "" {
		cfg, err := jaegercfg.FromEnv()
		if err != nil {
			// parsing errors might happen here, such as when we get a string where we expect a number
			log.Printf("Could not parse Jaeger env vars: %s", err.Error())
			return
		}

		metricsFactory := prometheus.New()
		tracer, closer, err := cfg.NewTracer(
			config.Metrics(metricsFactory),
		)
		if err != nil {
			log.Printf("Could not initialize jaeger tracer: %s", err.Error())
			return
		}
		defer closer.Close()

		opentracing.SetGlobalTracer(tracer)
	} else {
		fmt.Println("Jaeger host not set, skipping")
	}
	// TRACER DONE
}
