package errno

// code=0 请求成功
// code=4xx 客户端请求错误
// code=5xx 服务器端错误
// code=2xxxx 业务处理错误码

type Errno struct {
	Code    int
	Message string
}

func (e *Errno) Success() bool {
	return e.Code == OK.Code
}

// 全局通用错误码定义
// 2xx/4xx/5xx
var (
	OK = &Errno{Code: 200, Message: "Success"}

	ErrMissingParameter = &Errno{Code: 400, Message: "Missing parameter %s"}
	ErrParameterInvalid = &Errno{Code: 400, Message: "Invalid parameter %s"}
	ErrUnauthorized     = &Errno{Code: 401, Message: "Unauthorized"}
	ErrPermissionDeby   = &Errno{Code: 403, Message: "Permission deny"}
	ErrRouteNotFound    = &Errno{Code: 404, Message: "Route not found"}

	ErrInternalServer      = &Errno{Code: 500, Message: "Internal server error"}
	ErrInternalServerPanic = &Errno{Code: 500, Message: "Internal server panic occurred"}
	ErrDatabase            = &Errno{Code: 500, Message: "Database error"}
	ErrUnknown             = &Errno{Code: 510, Message: "Unknown error"}
)
