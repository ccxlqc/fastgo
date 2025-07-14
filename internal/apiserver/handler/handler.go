package handler

import (
	"github.com/onexstack/fastgo/internal/apiserver/biz"
	"github.com/onexstack/fastgo/internal/apiserver/pkg/conversion/validation"
)

type Handler struct {
	biz biz.IBiz
	val *validation.Validator
}

func NewHandler(biz biz.IBiz, val *validation.Validator) *Handler {
	return &Handler{
		biz: biz,
		val: val,
	}
}
