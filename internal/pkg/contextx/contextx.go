package contextx

import "context"

type (
	// requestIDKey 定义请求 ID 的上下文键.
	requestIDKey struct{}
	// userIDKey 定义用户 ID 的上下文键.
	userIDKey struct{}
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

func RequestID(ctx context.Context) string {
	// 类型断言，将 ctx 中的值转换为 string 类型
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}

// WithUserID 将用户 ID 存放到上下文中.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// UserID 从上下文中提取用户 ID.
func UserID(ctx context.Context) string {
	userID, _ := ctx.Value(userIDKey{}).(string)
	return userID
}
