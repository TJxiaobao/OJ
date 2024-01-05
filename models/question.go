package models

type Question struct {
	BaseModel
	QuestionId int    `gorm:"column:question_id" json:"-"`
	Input      string `gorm:"column:input" json:"-"`
	Output     string `gorm:"column:output" json:"-"`
	Title      string `gorm:"column:title" json:"-"`
	Content    string `gorm:"column:content" json:"-"`
	Tags       string `gorm:"column:tags" json:"-"`
	SubmitNum  int    `gorm:"column:submit_num" json:"-"`
	PassNum    int    `gorm:"column:pass_num" json:"-"`
	MaxRunTime int    `gorm:"column:max_runtime" json:"-"`
	MaxMem     int    `gorm:"column:max_mem" json:"-"`
}

func (table Question) TableName() string {
	return "question"
}
