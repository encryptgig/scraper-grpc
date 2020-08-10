package model

type Key struct {
	Model   `json:"model"`
	KeyByte []byte  `gorm:"unique;not null"`
	IV      []byte  `gorm:"unique;not null"`
	Algo    *string `gorm:"not null"`
	Size    *int    `gorm:"not null"`
	OwnerID *string `gorm:"not null"`
	Meta    string
}
