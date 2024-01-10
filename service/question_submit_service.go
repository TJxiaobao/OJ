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

func GetQuestionSubmitList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", constant.DefaultPage))
	if err != nil {
		log.Print("page error", err)
		restapi.Failed(c, err)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", constant.DefaultSize))
	if err != nil {
		log.Print("size", err)
		restapi.Failed(c, err)
		return
	}

	query := cqe.GetQuestionSubmitListQuery{}
	if err = c.ShouldBindJSON(query); err != nil {
		log.Print("question_submit query error", err)
		restapi.Failed(c, err)
		return
	}

	data, count, err := dao.GetQuestionSubmitList(query.UserId, query.QuestionId, query.Status, page, size)
	if err != nil {
		restapi.FailedWithStatus(c, err, 500)
		return
	}
	resp := restapi.NewPageResult(data, page, size, count)
	restapi.Success(c, resp)
}
