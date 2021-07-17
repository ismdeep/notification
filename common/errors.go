package common

// Error error string
type Error string

func (e Error) Error() string { return string(e) }

const (
	// ErrNotImplemented 未实现的功能
	ErrNotImplemented = Error("not implemented")

	// ErrBadRequest 请求参数错误
	ErrBadRequest = Error("请求参数错误")

	// ErrSystemError 系统错误
	ErrSystemError = Error("系统错误")

	// ErrUserNotExists 用户不存在
	ErrUserNotExists = Error("用户不存在")

	// ErrUserHasBeenBaned 账号被禁用
	ErrUserHasBeenBaned = Error("账号被禁用")

	// ErrNotLogin 尚未登录
	ErrNotLogin = Error("尚未登录")

	// ErrDatabaseOperateFailed 数据库操作失败
	ErrDatabaseOperateFailed = Error("数据库操作失败")
)
