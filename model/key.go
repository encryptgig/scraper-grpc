package model

import "github.com/jinzhu/gorm"

type Key struct {
	Model   `json:"model"`
	KeyByte []byte  `gorm:"unique;not null"`
	IV      []byte  `gorm:"unique;not null"`
	Algo    *string `gorm:"not null"`
	Size    *int    `gorm:"not null"`
	OwnerID *string `gorm:"not null"`
	Meta    string
}

func (k *Key) CreateKey(db *gorm.DB) error {
	return db.Create(k).Error
}

func (g *Key) ListKey(db *gorm.DB, owner string ) (*[]Key, error) {
	out := &[]Key{}
	err := db.Where("owner_id = ?",owner).Find(out).Error
	return out, err
}