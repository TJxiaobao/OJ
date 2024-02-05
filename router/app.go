package router

import (
	"github.com/TJxiaobao/OJ/docs"
	"github.com/TJxiaobao/OJ/service"
	"github.com/TJxiaobao/OJ/utils/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// question
	r.GET("/question/app", service.Test)
	r.GET("/question/get_question", service.GetQuestionList)
	r.GET("/question/get_question_detail", service.GetQuestionDetail)
	r.GET("/question/submit_list", service.GetQuestionSubmitList)
	r.POST("/question/submit", service.QuestionSubmit)

	// user
	r.GET("/user/get_user_detail", service.GetUserDetail)
	r.POST("/user/delete_user", service.DeleteUser)
	r.POST("/user/login", service.Login)
	r.POST("/user/register", service.Register)
	r.POST("/user/send_code", service.SendCode)

	return r
}
