package tracing

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
)

func GetSpan(ctx context.Context, spanName string) opentracing.Span {
	spanRaw := ctx.Value("SPAN")

	if spanRaw != nil {
		parentSpan := ctx.Value("SPAN").(opentracing.Span)

		sp := opentracing.StartSpan(
			spanName,
			opentracing.ChildOf(parentSpan.Context()))

		return sp
	}

	sp := opentracing.StartSpan(spanName)

	return sp
}
