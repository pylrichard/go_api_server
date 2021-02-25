package model

import (
	"fmt"

	"go/tiny_http_server/pkg/auth"
	"go/tiny_http_server/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
)

// UserModel 代表注册用户
type UserModel struct {
	BaseModel
	UserName string `json:"user_name" gorm:"column:user_name;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// TableName 获取表名
func (u *UserModel) TableName() string {
	return "user_info"
}

// Create 创建新用户
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser 根据id删除相应用户
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id

	return DB.Self.Delete(&user).Error
}

// Update 更新用户
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser 根据id获取用户
func GetUser(name string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("user_name = ?", name).First(&u)

	return u, d.Error
}

// ListUser 获取用户列表
func ListUser(name string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	users := make([]*UserModel, 0)
	var count uint64
	where := fmt.Sprintf("user_name like '%%%s%%'", name)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare 比较用户密码
func (u *UserModel) Compare(pwd string) error {
	return auth.Compare(u.Password, pwd)
}

// Encrypt 对用户密码进行加密
func (u *UserModel) Encrypt() error {
	u.Password, err := auth.Encrypt(u.Password)

	return err
}

// Validate 对字段进行验证
func (u *UserModel) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}