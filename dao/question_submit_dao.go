package dao

import (
	"github.com/TJxiaobao/OJ/models"
	"github.com/TJxiaobao/OJ/utils/logger_oj"
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
		tx = tx.Where("question_id = ?", question_id)
	}
	if status != 0 {
		tx = tx.Where("status = ?", status)
	}
	err := DB.Model(models.QuestionSubmit{}).
		Where(tx).
		Where("isDelete = 0").
		Count(&count).
		Offset(page - 1).
		Limit(size).
		Find(&data).
		Error
	if err != nil {
		log.Print("getQuestionSubmitDao error", err)
		logger_oj.LogToFile("getQuestionSubmitDao error: " + err.Error())
		return nil, 0, err
	}
	return data, count, nil
}

func SaveQuestionSubmitInfo(question_submit_info models.QuestionSubmit) error {
	return DB.Model(models.QuestionSubmit{}).Create(question_submit_info).Error
}
