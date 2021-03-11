package errno

var (
	// Common errors

	// Ok 正确
	Ok = &Errno{Code: 0, Msg: "Ok"}
	// InternalServerError 内部错误
	InternalServerError = &Errno{Code: 10001, Msg: "Internal server error."}
	// ErrBind 绑定错误
	ErrBind = &Errno{Code: 10002, Msg: "Error occurred while binding the request body to the struct."}

	// ErrValidation 验证错误
	ErrValidation = &Errno{Code: 20001, Msg: "Validation failed."}
	// ErrDatabase 操作数据库错误
	ErrDatabase = &Errno{Code: 20002, Msg: "Database error."}
	// ErrToken JWT错误
	ErrToken = &Errno{Code: 20003, Msg: "Error occurred while signing the JWT."}

	// User errors

	// ErrEncrypt 加密错误
	ErrEncrypt = &Errno{Code: 20101, Msg: "Error occurred while encrypting the password."}
	// ErrUserNotFound 没有找到对应用户信息
	ErrUserNotFound = &Errno{Code: 20102, Msg: "The user was not found."}
	// ErrTokenInvalid Token错误
	ErrTokenInvalid = &Errno{Code: 20103, Msg: "The token was invalid."}
	// ErrPasswordIncorrect 密码不正确
	ErrPasswordIncorrect = &Errno{Code: 20104, Msg: "The password was incorrect."}
)
