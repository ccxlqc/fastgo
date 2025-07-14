package model

import (
	"github.com/onexstack/fastgo/internal/pkg/rid"
	"github.com/onexstack/fastgo/pkg/auth"
	"gorm.io/gorm"
)

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (m *User) BeforeCreate(tx *gorm.DB) error {
	// Encrypt the user password.
	var err error
	m.Password, err = auth.Encrypt(m.Password)
	if err != nil {
		return err
	}

	return nil
}

func (m *Post) AfterCreate(tx *gorm.DB) (err error) {
	m.PostID = rid.PostID.New(uint64(m.ID))

	return tx.Save(m).Error
}

// AfterCreate 在创建数据库记录之后生成 userID.
func (m *User) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))

	return tx.Save(m).Error
}
