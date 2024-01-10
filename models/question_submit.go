package models

type QuestionSubmit struct {
	BaseModel
	QuestionSubmitId string `gorm:"column:question_submit_id" json:"-"`
	UserId           string `gorm:"column:user_id" json:"-"`
	QuestionId       string `gorm:"column:question_id" json:"-"`
	Language         int    `gorm:"column:language" json:"-"`
	Code             string `gorm:"column:code" json:"-"`
	JudgeInfo        string `gorm:"column:judge_info" json:"-"`
	Status           int    `gorm:"column:status" json:"-"`
}

func (table QuestionSubmit) TableName() string {
	return "oj_question_submit_info"
}
