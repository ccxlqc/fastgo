package biz

import (
	post "github.com/onexstack/fastgo/internal/apiserver/biz/v1/post"
	user "github.com/onexstack/fastgo/internal/apiserver/biz/v1/user"
	"github.com/onexstack/fastgo/internal/apiserver/store"
)

type IBiz interface {
	UserV1() user.UserBiz
	PostV1() post.PostBiz
}

type biz struct {
	store store.IStore
}

var _ IBiz = (*biz)(nil)

func New(store store.IStore) *biz {
	return &biz{
		store: store,
	}
}

func (b *biz) UserV1() user.UserBiz {
	return user.New(b.store)
}

func (b *biz) PostV1() post.PostBiz {
	return post.New(b.store)
}
