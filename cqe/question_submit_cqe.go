package cqe

import "github.com/TJxiaobao/OJ/utils/errno"

type GetQuestionSubmitListQuery struct {
	UserId     string `json:"user_id" form:"user_id"`
	QuestionId string `json:"question_id" form:"question_id"`
	Status     int    `json:"status" form:"status"`
}

type QuestionSubmitCmd struct {
	UserId     string `json:"user_id" form:"user_id"`
	QuestionId string `json:"question_id" form:"question_id"`
	Code       string `json:"code" form:"code"`
	Language   string `json:"language" form:"language"`
	Input      string `json:"input" form:"input"`
}

func (q *QuestionSubmitCmd) Validate() error {
	if q.Language == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "language")
	}
	if q.UserId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_id")
	}
	if q.QuestionId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "question_id")
	}
	if q.Code == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code")
	}
	return nil
}
