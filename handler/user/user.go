package user

// CreateRequest 创建用户请求
type CreateRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// CreateResponse 创建用户响应
type CreateResponse struct {
	UserName string `json:user_name`
}