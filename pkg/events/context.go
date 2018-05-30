package events

import "context"

type Context map[string]interface{}

type contextKeyType string

const contextKey contextKeyType = "context"

func WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey, map[string]interface{}{})
}

func GetContext(ctx context.Context) Context {
	v := ctx.Value(contextKey)
	if v == nil {
		return nil
	}

	return v.(Context)
}
