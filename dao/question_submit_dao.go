package dao

import (
	"github.com/TJxiaobao/OJ/models"
	"log"
)

func GetQuestionSubmitList(user_id, question_id string, status, page, size int) ([]*models.QuestionSubmit, int64, error) {
	var count int64
	data := make([]*models.QuestionSubmit, 0)
	tx := DB
	if user_id != "" {
		tx = tx.Where("user_id = ?", user_id)
	}
	if question_id != "" {
		tx = tx.Where("question = ?", question_id)
	}
	if status != -1 {
		tx = tx.Where("status = ?", status)
	}
	err := DB.Model(models.QuestionSubmit{}).
		Count(&count).
		Find(data).
		Where(tx).
		Offset(page).
		Limit(size).
		Where("isDelete = 0").
		Error
	if err != nil {
		log.Print("getQuestionSubmitDao error", err)
		return nil, 0, err
	}
	return data, count, nil
}

func SaveQuestionSubmitInfo(question_submit_info models.QuestionSubmit) error {
	return DB.Model(models.QuestionSubmit{}).Create(question_submit_info).Error
}
