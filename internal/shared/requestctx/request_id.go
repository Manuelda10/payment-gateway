package requestctx

import "context"

type keyRequestID struct{}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, keyRequestID{}, id)
}

func RequestID(ctx context.Context) (string, bool) {
	v := ctx.Value(keyRequestID{})
	s, ok := v.(string)
	return s, ok
}
