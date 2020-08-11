package model

import "github.com/jinzhu/gorm"
const (
	PermissionRead = "PermissionReadKey"
	PermissionUpdate = "PermissionUpdateKey"
	PermissionAdd = "PermissionAddKey"
	PermissionDelete = "PermissionDeleteKey"
)

type Model struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}


type Permission struct {
	Model
	ResourceID 	string `gorm:"primary_key"`
	GroupID    	string `gorm:"primary_key"`
	Permission 	string `gorm:"primary_key"`
}

func (g *Permission) CreatePermission(db *gorm.DB) error {
	return db.Create(g).Error
}

func (g *Permission) ListGroupForPermission( db *gorm.DB, perms []string  ) (*[]Permission , error) {
	out := &[]Permission{}
	err := db.Select("distinct(group_id)").Where("Permission IN (?)", perms).Find(out).Error
	return out, err
}

func (g *Permission) ListResourcesForGroupAndPermission( db *gorm.DB, perms []string, groups []string  ) (*[]Permission , error) {
	out := &[]Permission{}
	err := db.Select("resource_id").Where("permission IN (?)", perms).Where("group_id IN (?)", groups ).Find(out).Error
	return out, err
}

