package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/onexstack/fastgo/internal/pkg/core"
	v1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
	"github.com/onexstack/onexstack/pkg/errorsx"
)

func (h *Handler) CreateUser(c *gin.Context) {
	slog.Info("Create user function called")

	var rq v1.CreateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateCreateUserRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Create(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	slog.Info("Update user function called")

	var rq v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateUpdateUserRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Update(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	slog.Info("Delete user function called")

	var rq v1.DeleteUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateDeleteUserRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Delete(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) GetUser(c *gin.Context) {
	slog.Info("Get user function called")

	var rq v1.GetUserRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateGetUserRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.UserV1().Get(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) ListUser(c *gin.Context) {
	slog.Info("List user function called")

	var rq v1.ListUserRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateListUserRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.UserV1().List(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}
