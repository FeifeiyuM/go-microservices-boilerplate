package gmberror

// 定义
type GMBError interface {
	// http 状态码
	HttpStatus() int
	// 错误唯一编号
	Code() string
	// 是否需要监控告警
	SendMsg() bool
	// 错误描述，这个应该是对用户友好的错误提示
	Message() string
	// 存放原始错误，错误堆栈
	Error() string
}

type errInfo struct {
	httpStatus int    // http status
	code       string // 自定义错误code
	message      string // 错误描述
	sendMsg  bool   // 是否发送邮件
	err      error  // 错误信息
}
// HttpStatus 获取http 状态值
func (e *errInfo) HttpStatus() int {
	return e.httpStatus
}
// Code 获取自定义 code
func (e *errInfo) Code() string {
	return e.code
}
// SendMail 是否发送告警
func (e *errInfo) SendMsg() bool {
	return e.sendMsg
}
// Alert 获取错描述
func (e *errInfo) Message() string {
	return e.message
}
// Error 错误详情或堆栈
func (e *errInfo) Error() string {
	return e.Error()
}
// NewGMBError 新建错误对象
func NewGMBError(httpCode int, code, message string, sendMsg bool, err error) GMBError {
	return &errInfo{
		httpStatus: httpCode,
		code:       code,
		message:    message,
		sendMsg:    sendMsg,
		err:        err,
	}
}

