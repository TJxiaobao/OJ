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

	query := cqe.GetQuestionQuery{}
	if err = c.ShouldBindQuery(&query); err != nil {
		log.Print("question query error", err)
		restapi.Failed(c, err)
		return
	}

	data, count := dao.GetQuestionList(query.Keyword, page, size)
	resp := restapi.NewPageResult(data, page, size, count)
	restapi.Success(c, resp)
}

// GetQuestionDetail godoc
// @Summary      查询Question 详情
// @Description	 test
// @Tags         question
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Question ID"
// @Success      200  {string} json "{"code" : "200", "msg" : "", "data" : ""}"
// @Router       /question/{id} [get]

func GetQuestionDetail(c *gin.Context) {
	query := cqe.GetQuestionDetailQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Print("question query error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 question_id 是否为空
	if err := query.Validate(); err != nil {
		log.Print("question_id must not null", err)
		restapi.Failed(c, err)
		return
	}

	data, err := dao.GetQuestionDetail(query.QuestionId)
	if data.QuestionId == "" {
		restapi.Success(c, "This problem does not exist, please contact the staff！")
		return
	}
	if err != nil {
		restapi.Failed(c, err)
		return
	}
	restapi.Success(c, data)
}
