package models

type Question struct {
	BaseModel
	QuestionId string `gorm:"column:question_id" json:"question_id"`
	Input      string `gorm:"column:input" json:"input"`
	Output     string `gorm:"column:output" json:"output"`
	Title      string `gorm:"column:title" json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	Tags       string `gorm:"column:tags" json:"tags"`
	SubmitNum  int    `gorm:"column:submit_num" json:"submit_num"`
	PassNum    int    `gorm:"column:pass_num" json:"pass_num"`
	MaxRunTime int    `gorm:"column:max_runtime" json:"max_run_time"`
	MaxMem     int    `gorm:"column:max_mem" json:"max_mem"`
	Difficulty int    `gorm:"column:difficulty" json:"difficulty"`
}

func (table Question) TableName() string {
	return "oj_question"
}
