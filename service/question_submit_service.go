package service

import (
	"encoding/json"
	"github.com/TJxiaobao/OJ/constant"
	"github.com/TJxiaobao/OJ/cqe"
	"github.com/TJxiaobao/OJ/dao"
	"github.com/TJxiaobao/OJ/judge/codeExecute"
	"github.com/TJxiaobao/OJ/models"
	"github.com/TJxiaobao/OJ/utils/restapi"
	"github.com/TJxiaobao/OJ/utils/uuid"
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
	if err = c.ShouldBindQuery(&query); err != nil {
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

func QuestionSubmit(c *gin.Context) {
	cmd := cqe.QuestionSubmitCmd{}
	if err := c.BindJSON(&cmd); err != nil {
		log.Fatal("Parameter error:", err)
		restapi.Failed(c, err)
		return
	}

	// 执行代码
	res := codeExecute.CodeExecute(cmd.Language, cmd.Code, cmd.Input)

	// 刷库
	judgeInfo, err := json.Marshal(res.JudgeInfo)
	if err != nil {
		log.Fatal("json error:", err)
		restapi.Failed(c, err)
		return
	}
	questionSubmit := models.QuestionSubmit{
		QuestionSubmitId: uuid.GetUUID(),
		UserId:           cmd.UserId,
		QuestionId:       cmd.QuestionId,
		Language:         cmd.Language,
		Code:             cmd.Code,
		JudgeInfo:        string(judgeInfo),
		Status:           res.Status,
	}
	err = dao.SaveQuestionSubmitInfo(questionSubmit)
	if err != nil {
		log.Fatal("DB error:", err)
		restapi.Failed(c, err)
		return
	}
	restapi.Success(c, res)
}
