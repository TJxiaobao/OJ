package constant

var (
	DefaultPage = "1"
	DefaultSize = "1"
	SandBoXUrl  = "http://localhost:8081/code_execute"

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

	// email constant
	EmailSender    = "test@qq.com"
	EmailAuthCode  = "auth"
	EmailTitl      = "OJ 验证码"
	EmailBodyTem   = "您的注册验证码为："
	EmailQQHost    = "smtp.qq.com"
	Email163Host   = "smtp.163.com"
	EmailSinaHost  = "smtp.sina.com"
	EmailYahooHost = "smtp.yahoo.com"
	Email126Host   = "smtp.126.com"
	EmailSohuHost  = "smtp.sohu.com"

	// phone constant
)
