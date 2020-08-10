package model

import "github.com/jinzhu/gorm"

type Group struct {
	Name   string `gorm:"primary_key"`
	UserID string `gorm:"primary_key"`
}

func NewGroup(name string, userid string) *Group {
	if len(name) == 0 || len(userid) == 0 {
		return nil
	}
	return &Group{Name: name, UserID: userid}
}

func (g *Group) CreateGroup(db *gorm.DB) error {
	return db.Create(g).Error
}

func (g *Group) ListGroups(db *gorm.DB) (*[]Group, error) {
	out := &[]Group{}
	err := db.Select("distinct(name)").Find(out).Error
	return out, err
}

func (g *Group) ListGroupsForUser(db *gorm.DB, userid string) (*[]Group, error) {
	out := &[]Group{}
	err := db.Where("user_id = ?", userid).Find(out).Error
	return out, err
}

func (g *Group) ListUserForGroups(db *gorm.DB, group string) (*[]Group, error) {
	out := &[]Group{}
	err := db.Where("name = ?", group).Find(out).Error
	return out, err
}
