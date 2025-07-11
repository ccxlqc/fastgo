package contextx

import "context"

type (
	requestIDKey struct{}
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

func RequestID(ctx context.Context) string {
	// 类型断言，将 ctx 中的值转换为 string 类型
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}
