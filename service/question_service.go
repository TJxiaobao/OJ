package service

import (
	"OJ/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetQuestionList godoc
// @Summary      查询Question问题列表
// @Description	 test
// @Tags         question
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Question ID"
// @Success      200  {string} json "{"code" : "200", "msg" : "", "data" : ""}"
// @Router       /question/{id} [get]
func GetQuestionList(c *gin.Context) {
	dao.GetQuestionList()
	c.String(http.StatusOK, "success")
}
