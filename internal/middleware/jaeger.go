package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"frame/pkg/cfg"
	"frame/pkg/jaeger"
)

func Jaeger() func(c *gin.Context) {
	return func(c *gin.Context) {
		if cfg.Viper().GetBool("jaeger.open") {
			var parentSpan opentracing.Span

			tracer, closer := jaeger.NewJaegerTracer(
				cfg.Viper().GetString("app.name"),
				fmt.Sprintf("%s:%d",
					cfg.Viper().GetString("jaeger.host"),
					cfg.Viper().GetInt("jaeger.port"),
				))
			defer closer.Close()

			spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			if err != nil {
				parentSpan = tracer.StartSpan(c.Request.URL.Path)
				defer parentSpan.Finish()
			} else {
				parentSpan = opentracing.StartSpan(
					c.Request.URL.Path,
					opentracing.ChildOf(spCtx),
					opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
					ext.SpanKindRPCServer,
				)
				defer parentSpan.Finish()
			}

			c.Set("Tracer", tracer)
			c.Set("ParentSpanContext", parentSpan.Context())
		}

		c.Next()
	}
}
