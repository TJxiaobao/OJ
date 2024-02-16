package service

import (
	"github.com/TJxiaobao/OJ/utils/logger"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	logger.LogToFile("test log file")
	c.JSON(200, "test")
}
