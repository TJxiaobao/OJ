package service

import (
	"OJ/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuestionList(c *gin.Context) {
	dao.GetQuestionList()
	c.String(http.StatusOK, "success")
}
