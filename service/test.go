package service

import (
	"github.com/TJxiaobao/OJ/utils/logger_oj"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	logger_oj.LogToFile("test log file")
	c.JSON(200, "test")
}
