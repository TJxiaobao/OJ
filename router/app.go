package router

import (
	"OJ/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/app", service.Test)
	r.GET("/get_question", service.GetQuestionList)

	return r
}
