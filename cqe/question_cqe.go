package cqe

import "github.com/TJxiaobao/OJ/utils/errno"

// cmd query event

type GetQuestionQuery struct {
	Keyword string `json:"keyword"`
	Type    string `json:"type"`
}

type GetQuestionDetailQuery struct {
	QuestionId string `json:"question_id"`
}

func (q *GetQuestionDetailQuery) Validate() error {
	if q.QuestionId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "question_id")
	}
	return nil
}
