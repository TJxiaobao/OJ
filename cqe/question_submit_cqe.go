package cqe

type GetQuestionSubmitListQuery struct {
	UserId     string `json:"user_id"`
	QuestionId string `json:"question_id"`
	Status     int    `json:"status"`
}
