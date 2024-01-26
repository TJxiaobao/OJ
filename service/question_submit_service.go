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

	// 判断是否存在改题目
	que, err := dao.GetQuestionDetail(cmd.QuestionId)
	if que == nil {
		restapi.Success(c, "没有该题目!")
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

	// 业务处理
	question, err := dao.GetQuestionDetail(cmd.QuestionId)

	// 1、 然后判断是否超时或者超出内存
	if res.Status == constant.StatusOK && res.JudgeInfo.Time > question.MaxRunTime {
		log.Print("outTime ", question, cmd.UserId)
		restapi.Success(c, "超时！")
		return
	}
	if res.Status == constant.StatusOK && res.JudgeInfo.Memory > question.MaxMem {
		log.Print("out max mem ", question, cmd.UserId)
		restapi.Success(c, "内存过大！")
		return
	}

	// 2、 判断输出是否正确
	if res.Status == constant.StatusOK && res.Output == question.Output {
		restapi.Success(c, "AC!")
		return
	} else if res.Status == constant.StatusOK && res.Output != question.Output {
		restapi.Success(c, "输出错误！")
		return
	}
	// 3、 首先判断 status 如果不是执行成功，则返回。
	restapi.Success(c, res)
}
