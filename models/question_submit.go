package models

type QuestionSubmit struct {
	BaseModel
	QuestionSubmitId string `gorm:"column:question_submit_id" json:"-"`
	UserId           string `gorm:"column:user_id" json:"user_id"`
	QuestionId       string `gorm:"column:question_id" json:"question_id"`
	Language         string `gorm:"column:language" json:"language"`
	Code             string `gorm:"column:code" json:"code"`
	JudgeInfo        string `gorm:"column:judge_info" json:"judge_info"`
	Status           int    `gorm:"column:status" json:"status"`
}

func (table QuestionSubmit) TableName() string {
	return "oj_question_submit_info"
}
