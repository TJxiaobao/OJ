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
		Count(&count).
		Where("title LIKE ? ", "%"+keyword+"%").
		Where("content LIKE ? ", "%"+keyword+"%").
		Offset(page).
		Limit(size).
		Find(&data).Error
	if err != nil {
		log.Print("GetQuestionList error", err)
		return nil, 0
	}
	return data, count
}
