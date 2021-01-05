package tracing

import (
	"context"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
)

func JaegerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		wireContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))

		if err == nil {
			ctx = context.WithValue(ctx, "span", wireContext)
		}

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
