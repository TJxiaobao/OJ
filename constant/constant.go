package constant

import config2 "github.com/TJxiaobao/OJ/utils/config"

var config = config2.LoadConfig()

var (
	// todo 后期改为json配置文件，并读取配置文件

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
	EmailSender   = config.Email.Sender
	EmailAuthCode = config.Email.AuthCode
	EmailTitl     = "OJ 验证码"
	EmailBodyTem  = "您的注册验证码为："
	EmailHost     = config.Email.Host
	EmailPort     = config.Email.Port

	// phone constant

	// GitHub oauth
	GithubOauthClientId     = ""
	GithubOauthClientSecret = ""
	GithubOauthRedirectUrl  = ""
)
