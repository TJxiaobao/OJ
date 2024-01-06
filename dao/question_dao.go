package dao

import (
	"OJ/models"
	"log"
)

func GetQuestionList() {
	data := make([]*models.Question, 0)
	err := DB.Find(&data).Error
	if err != nil {
		log.Print("GetQuestionList error")
	}

}
