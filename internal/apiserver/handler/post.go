package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/onexstack/fastgo/internal/pkg/core"
	v1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
	"github.com/onexstack/onexstack/pkg/errorsx"
)

func (h *Handler) CreatePost(c *gin.Context) {
	slog.Info("Create post function called")

	var rq v1.CreatePostRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)

		return
	}

	if err := h.val.ValidateCreatePostRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Create(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	slog.Info("Update post function called")

	var rq v1.UpdatePostRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateUpdatePostRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Update(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) DeletePost(c *gin.Context) {
	slog.Info("Delete post function called")

	var rq v1.DeletePostRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	resp, err := h.biz.PostV1().Delete(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) GetPost(c *gin.Context) {
	slog.Info("Get post function called")

	var rq v1.GetPostRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateGetPostRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.PostV1().Get(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

func (h *Handler) ListPost(c *gin.Context) {
	slog.Info("List post function called")

	var rq v1.ListPostRequest
	if err := c.ShouldBindQuery(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateListPostRequest(c, &rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrInvalidArgument.WithMessage("%s", err.Error()))
		return
	}

	resp, err := h.biz.PostV1().List(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}
