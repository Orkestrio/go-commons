package cache

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	opentracing "github.com/opentracing/opentracing-go"
	ext "github.com/opentracing/opentracing-go/ext"
)

type ResponseCache struct {
	Cache graphql.Cache
}

var _ interface {
	graphql.ResponseInterceptor
	graphql.HandlerExtension
} = ResponseCache{}

func (a ResponseCache) ExtensionName() string {
	return "ResponseCache"
}

func (a ResponseCache) Validate(schema graphql.ExecutableSchema) error {
	if a.Cache == nil {
		return fmt.Errorf("ResponseCache.Cache can not be nil")
	}
	return nil
}

func (a ResponseCache) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	rc := graphql.GetOperationContext(ctx)

	if (rc.Operation == nil) || (rc.Operation.Operation == "mutation") {
		return next(ctx)
	}

	spanRaw := ctx.Value("span")
	sp := opentracing.StartSpan("cache")

	if spanRaw != nil {
		wireContext := spanRaw.(opentracing.SpanContext)
		sp = opentracing.StartSpan(
			"cache",
			ext.RPCServerOption(wireContext))
	}

	query := rc.RawQuery
	queryHash := computeQueryHash(fmt.Sprintf("%s:%v", query, rc.Variables))
	value, found := a.Cache.Get(ctx, queryHash)

	if !found {
		// Cache miss
		resp := next(ctx)
		respCopy := *resp

		fmt.Println("Request: ", query)
		fmt.Println("Response:", string(respCopy.Data), ", errors: ", respCopy.Errors)

		if len(respCopy.Errors) == 0 {
			a.Cache.Add(ctx, queryHash, respCopy)
		}

		return resp
	}

	valueTyped := value.(graphql.Response)
	fmt.Println("Request from cache: ", query)
	fmt.Println("Response from cache:", string(valueTyped.Data), ", errors: ", valueTyped.Errors)

	if rc.Operation.Name != "IntrospectionQuery" {
		sp.Finish()
	}

	return &valueTyped
}
