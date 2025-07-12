package model

import (
	"github.com/onexstack/fastgo/internal/pkg/rid"
	"gorm.io/gorm"
)

func (m *Post) AfterCreate(tx *gorm.DB) (err error) {
	m.PostID = rid.PostID.New(uint64(m.ID))

	return tx.Save(m).Error
}

// AfterCreate 在创建数据库记录之后生成 userID.
func (m *User) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))

	return tx.Save(m).Error
}
