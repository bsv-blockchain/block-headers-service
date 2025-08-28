package metrics

import (
	"context"

	"github.com/bitcoin-sv/block-headers-service/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	otelgrpc "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer

func EnableTracing(log *zerolog.Logger, cfg *config.TracingConfig) (func(), error) {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("Block Headers Service"),
		semconv.ServiceVersionKey.String(config.Version()),
	)

	exporter, err := otelgrpc.New(
		context.Background(),
		otelgrpc.WithEndpoint(cfg.Endpoint),
		otelgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)

	Tracer = tp.Tracer("BHS")

	cleanup := func() {
		tlog := log.With().Str("subservice", "tracing").Logger()
		err = exporter.Shutdown(context.Background())
		if err != nil {
			tlog.Error().Err(err).Msg("Failed to shutdown tracing exporter")
		}

		err = tp.Shutdown(context.Background())
		if err != nil {
			tlog.Error().Err(err).Msg("Failed to shutdown tracing provider")
		}
	}

	return cleanup, nil
}

func RegisterTracing(gin *gin.Engine) {
	if Tracer != nil {
		gin.Use(otelgin.Middleware("block-headers-service"))
	}
}
