package user

import (
	"go/tiny_http_server/model"
)

// CreateRequest 创建用户请求
type CreateRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// CreateResponse 创建用户响应
type CreateResponse struct {
	UserName string `json:user_name`
}

// ListRequest 列表请求
type ListRequest struct {
	UserName string `json:"user_name"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

// ListResponse 列表响应
type ListResponse struct {
	Count    uint64            `json:"count"`
	UserList []*model.UserInfo `json:"user_list"`
}