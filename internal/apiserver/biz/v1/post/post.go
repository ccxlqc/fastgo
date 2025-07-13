package post

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/onexstack/fastgo/internal/apiserver/model"
	"github.com/onexstack/fastgo/internal/apiserver/pkg/conversion"
	"github.com/onexstack/fastgo/internal/apiserver/store"
	"github.com/onexstack/fastgo/internal/pkg/contextx"
	apiv1 "github.com/onexstack/fastgo/pkg/api/apiserver/v1"
	"github.com/onexstack/onexstack/pkg/store/where"
)

// PostBiz 定义处理帖子请求所需的方法.
type PostBiz interface {
	Create(ctx context.Context, rq *apiv1.CreatePostRequest) (*apiv1.CreatePostResponse, error)
	Update(ctx context.Context, rq *apiv1.UpdatePostRequest) (*apiv1.UpdatePostResponse, error)
	Delete(ctx context.Context, rq *apiv1.DeletePostRequest) (*apiv1.DeletePostResponse, error)
	Get(ctx context.Context, rq *apiv1.GetPostRequest) (*apiv1.GetPostResponse, error)
	List(ctx context.Context, rq *apiv1.ListPostRequest) (*apiv1.ListPostResponse, error)

	PostExpansion
}

// PostExpansion 定义额外的帖子操作方法.
type PostExpansion interface{}

type postBiz struct {
	store store.IStore
}

var _ PostBiz = (*postBiz)(nil)

func New(store store.IStore) *postBiz {
	return &postBiz{
		store: store,
	}
}

func (b *postBiz) Create(ctx context.Context, rq *apiv1.CreatePostRequest) (*apiv1.CreatePostResponse, error) {
	var postM model.Post
	_ = copier.Copy(&postM, rq)

	if err := b.store.Post().Create(ctx, &postM); err != nil {
		return nil, err
	}

	return &apiv1.CreatePostResponse{
		PostID: postM.PostID,
	}, nil
}

func (b *postBiz) Update(ctx context.Context, rq *apiv1.UpdatePostRequest) (*apiv1.UpdatePostResponse, error) {
	postM, err := b.store.Post().Get(ctx, where.F("userID", contextx.UserID(ctx), "postID", rq.PostID))
	if err != nil {
		return nil, err
	}

	if rq.Title != nil {
		postM.Title = *rq.Title
	}

	if rq.Content != nil {
		postM.Content = *rq.Content
	}

	if err := b.store.Post().Update(ctx, postM); err != nil {
		return nil, err
	}

	return &apiv1.UpdatePostResponse{}, nil
}

func (b *postBiz) Delete(ctx context.Context, rq *apiv1.DeletePostRequest) (*apiv1.DeletePostResponse, error) {
	whr := where.F("userID", contextx.UserID(ctx), "postID", rq.PostIDs)

	if err := b.store.Post().Delete(ctx, whr); err != nil {
		return nil, err
	}

	return &apiv1.DeletePostResponse{}, nil
}

func (b *postBiz) Get(ctx context.Context, rq *apiv1.GetPostRequest) (*apiv1.GetPostResponse, error) {
	postM, err := b.store.Post().Get(ctx, where.F("userID", contextx.UserID(ctx), "postID", rq.PostID))
	if err != nil {
		return nil, err
	}

	return &apiv1.GetPostResponse{
		Post: conversion.PostodelToPostV1(postM),
	}, nil
}

func (b *postBiz) List(ctx context.Context, rq *apiv1.ListPostRequest) (*apiv1.ListPostResponse, error) {
	whr := where.F("userID", contextx.UserID(ctx)).P(int(rq.Offset), int(rq.Limit))

	if rq.Title != nil {
		whr = whr.Q("title like ?", "%"+*rq.Title+"%")
	}

	count, postList, err := b.store.Post().List(ctx, whr)

	if err != nil {
		return nil, err
	}

	posts := make([]*apiv1.Post, 0, len(postList))
	for _, item := range postList {
		posts = append(posts, conversion.PostodelToPostV1(item))
	}

	return &apiv1.ListPostResponse{
		TotalCount: count,
		Posts:      posts,
	}, nil
}
