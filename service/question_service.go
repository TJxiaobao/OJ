package service

import (
	"github.com/TJxiaobao/OJ/constant"
	"github.com/TJxiaobao/OJ/cqe"
	"github.com/TJxiaobao/OJ/dao"
	"github.com/TJxiaobao/OJ/utils/restapi"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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
	page, err := strconv.Atoi(c.DefaultQuery("page", constant.DefaultPage))
	if err != nil {
		log.Print("page", err)
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", constant.DefaultSize))
	if err != nil {
		log.Print("size", err)
		return
	}
	query := cqe.GetQuestionQuery{}
	if err = c.ShouldBindJSON(query); err != nil {
		log.Print("question query error", err)
		return
	}

	data, count := dao.GetQuestionList(query.Keyword, page, size)
	resp := restapi.NewPageResult(data, page, size, count)
	restapi.Success(c, resp)
}
