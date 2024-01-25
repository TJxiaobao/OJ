package models

type Response struct {
	Output    string    `json:"output" form:"output"`   // 输出
	Message   string    `json:"message" form:"message"` // 接口信息
	Status    int       `json:"status" form:"form"`
	JudgeInfo JudgeInfo `json:"judge_info" form:"judge_info"`
}
