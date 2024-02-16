package router

import (
	"github.com/TJxiaobao/OJ/service"
	"github.com/TJxiaobao/OJ/utils/middlewares"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	// Test
	r.GET("/test", service.Test)

	// question
	r.GET("/question/app", service.Test)
	r.GET("/question/get_question", service.GetQuestionList)
	r.GET("/question/get_question_detail", service.GetQuestionDetail)
	r.GET("/question/submit_list", service.GetQuestionSubmitList)
	r.POST("/question/submit", service.QuestionSubmit)

	// user
	r.GET("/user/get_user_detail", service.GetUserDetail)
	r.GET("/user/login_github_url", service.LoginGetGithubUrl)
	r.GET("/user/login_github", service.LoginGitHub)
	r.POST("/user/delete_user", service.DeleteUser)
	r.POST("/user/login", service.Login)
	r.POST("/user/register", service.Register)
	r.POST("/user/send_code_email", service.SendCodeByEmail)
	r.POST("/user/send_code_sms", service.SendCodeBySms)
	r.POST("/user/login_sms", service.LoginSms)
	r.POST("/user/login_email", service.LoginEmail)

	return r
}
