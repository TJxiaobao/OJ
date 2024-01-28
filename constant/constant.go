package constant

var (
	DefaultPage = "1"
	DefaultSize = "1"
	SandBoXUrl  = "http://localhost:8081//code_execute"

	// 鉴权定义
	AUTH_REQUEST_HEADER = "auth"
	AUTH_REQUEST_SECRET = "test"

	// status定义
	StatusOK        = 1
	StatusOutTime   = 2
	StatusOutMemory = 3

	// 用户角色
	UserRoleSuperAdmin = "0"
	UserRoleAdmin      = "1"
	UserNormalRole     = "2"
)
