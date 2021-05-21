package gmberror

import "net/http"

// InvalidRequest 请求参数错误
func InvalidRequest(err error) GMBError {
	return NewGMBError(http.StatusBadRequest, "40001", "参数错误", false, err)
}

// ServerError 服务内部错误
func ServerError(err error) GMBError {
	return NewGMBError(http.StatusInternalServerError, "50001", "服务器开小差了,请稍后再试", true, err)
}

// 数据库异常
func DBError(err error) GMBError {
	return NewGMBError(http.StatusInternalServerError, "50002", "数据库异常", true, err)
}

// 无效的账户
func InvalidAccount(err error) GMBError {
	return NewGMBError(http.StatusBadRequest, "40002", "无效账户", false, err)
}

func InvalidOrder(err error) GMBError {
	return NewGMBError(http.StatusBadRequest, "40003", "无效订单", false, err)
}
