package models

type Request struct {
	Language string `json:"language" form:"language"`
	Code     string `json:"code" form:"code"`
	Input    string `json:"input" form:"input"`
}
