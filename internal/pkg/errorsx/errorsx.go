package errorsx

import (
	"errors"
	"fmt"
)

// ErrorX 定义了 fastgo 项目中使用的错误类型，用于描述错误的详细信息.
type ErrorX struct {
	// Code 表示错误的 HTTP 状态码，用于与客户端进行交互时标识错误的类型.
	Code int `json:"code,omitempty"`

	// Reason 表示错误发生的原因，通常为业务错误码，用于精准定位问题.
	Reason string `json:"reason,omitempty"`

	// Message 表示简短的错误信息，通常可直接暴露给用户查看.
	Message string `json:"message,omitempty"`
}

// New 创建一个新的错误.
func New(code int, reason string, format string, args ...any) *ErrorX {
	return &ErrorX{
		Code:    code,
		Reason:  reason,
		Message: fmt.Sprintf(format, args...),
	}
}

func (e *ErrorX) Error() string {
	return fmt.Sprintf("code: %d, reason: %s, message: %s", e.Code, e.Reason, e.Message)
}

func (e *ErrorX) WithMessage(format string, args ...any) *ErrorX {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func FromError(err error) *ErrorX {
	if err == nil {
		return nil
	}

	if errx := new(ErrorX); errors.As(err, &errx) {
		return errx
	}

	// 默认返回未知错误错误. 该错误代表服务端出错
	return New(ErrInternal.Code, ErrInternal.Reason, err.Error())
}
