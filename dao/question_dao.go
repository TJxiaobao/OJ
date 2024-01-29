package dao

import (
	"github.com/TJxiaobao/OJ/models"
	"log"
)

func GetQuestionList(keyword string, page int, size int) ([]*models.Question, int64) {
	data := make([]*models.Question, 0)
	var count int64
	err := DB.
		Model(models.Question{}).
		Where("title LIKE ? or content LIKE ? ", "%"+keyword+"%", "%"+keyword+"%").
		Where("isDelete = 0").
		Count(&count).
		Offset(page - 1).
		Limit(size).
		Find(&data).
		Error
	if err != nil {
		log.Print("GetQuestionList error", err)
		return nil, 0
	}
	return data, count
}

func GetQuestionDetail(question_id string) (*models.Question, error) {
	var data models.Question
	err := DB.Model(models.Question{}).Where("question_id = ? AND isDelete = 0", question_id).Find(&data).Error
	if err != nil {
		log.Print("question detail DB error", err)
		return nil, err
	}
	return &data, nil
}
