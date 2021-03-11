package errno

import "fmt"

// Errno 定制错误码
type Errno struct {
	Code	int
	Msg		string
}

// Error 返回错误码信息
func (err Errno) Error() string {
	return err.Msg
}

// Err 错误
type Err struct {
	Code	int
	Msg		string
	Err		error
}

// New 新建错误
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Msg: errno.Msg, Err: err}
}

// Add 为错误添加信息
func (err *Err) Add(msg string) Err {
	err.Msg += " " + msg

	return err
}

// Addf 为错误添加格式化信息
func (err *Err) Addf(format string, args ...interface{}) Err {
	err.Msg += " " + fmt.Sprintf(format, args...)

	return err
}

// Error 返回错误格式化信息
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, msg: %s, error: %s", err.Code, err.Msg, err.Err)
}

// IsErrUserNotFound 判断是否属于ErrUserNotFound类型错误
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)

	return code == ErrUserNotFound.Code
}

// DecodeErr 解析错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Ok.Code, Ok.Msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Msg
	case *Errno:
		return typed.Code, typed.Msg
	default:
	}
	
	return InternalServerError.Code, err.Error()
}