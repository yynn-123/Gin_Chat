package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Desc    string
	Type    string //消息类型:群聊、私聊、广播
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
