package model

type User struct {
	Model
	Domain   *string `gorm:"not null"`
	Name     *string `gorm:"not null"`
	Password *string `gorm:"not null"`
}
