package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onexstack/fastgo/internal/pkg/contextx"
	"github.com/onexstack/fastgo/internal/pkg/known"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 `x-request-id`，如果不存在则生成新的 UUID
		requestID := c.Request.Header.Get(known.XRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将 RequestID 保存到 context.Context 中，以便后续程序使用
		ctx := contextx.WithRequestID(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)

		// 将 RequestID 保存到 HTTP 返回头中，Header 的键为 `x-request-id`
		c.Writer.Header().Set(known.XRequestID, requestID)
		slog.Info("requestID", "requestID", requestID, "contextRequestID", contextx.RequestID(c.Request.Context()), "headerRequestID", c.Request.Header.Get(known.XRequestID))
		c.Next()
	}
}
