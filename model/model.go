package model

import (
	"sync"
	"time"
)

// BaseModel 基础对象模型
type BaseModel struct {
	Id          uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedTime time.Time  `gorm:"column:created_time" json:"-"`
	UpdatedTime time.Time  `gorm:"column:updated_time" json:"-"`
	DeletedTime *time.Time `gorm:"column:deleted_time" sql:"index" json:"-"`
}

// UserInfo 用户信息
type UserInfo struct {
	Id          uint64 `json:"id"`
	UserName    string `json:"user_name"`
	SayHello    string `json:"say_hello"`
	Password    string `json:"password"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

// UserList 用户信息集
type UserList struct {
	Lock  *sync.Mutex
	Infos map[uint64]*UserInfo
}

// Token JWT
type Token struct {
	Token string `json:"token"`
}