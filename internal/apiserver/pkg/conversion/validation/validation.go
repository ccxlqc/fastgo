package validation

import (
	"github.com/onexstack/fastgo/internal/apiserver/store"
)

type Validator struct {
	store *store.IStore
}

func NewValidator(store store.IStore) *Validator {
	return &Validator{store: &store}
}
