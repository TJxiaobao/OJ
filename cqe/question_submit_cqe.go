package cqe

type GetQuestionSubmitListQuery struct {
	UserId     string `json:"user_id" form:"user_id"`
	QuestionId string `json:"question_id" form:"question_id"`
	Status     int    `json:"status" form:"status"`
}
