package errno

var (
	/*
		Common erros
	*/

	// Ok 正确
	Ok					= &Errno{Code: 0, Msg: "Ok"}
	// InternalServerError 内部错误
	InternalServerError	= &Errno{Code: 10001, Msg: "Internal server error."}
	// ErrBind 绑定错误
	ErrBind				= &Errno{Code: 10002, Msg: "Error occurred while binding the request body to the struct."}

	/*
		User errors
	*/

	// ErrUserNotFound 没有找到对应用户信息
	ErrUserNotFound = &Errno{Code: 20102, Msg: "The user was not found."}
)