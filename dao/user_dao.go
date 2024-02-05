package dao

import (
	"errors"
	"github.com/TJxiaobao/OJ/models"
	"gorm.io/gorm"
	"log"
)

func GetUserDetail(user_id string) (*models.User, error) {
	var user models.User
	err := DB.Model(models.User{}).
		Where("user_id = ? and isDelete = 0", user_id).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 判断是否为记录不存在的错误
			return nil, nil // 返回空的 user
		}
		log.Print("db getUserDetail error ", err)
		return nil, err
	}
	return &user, nil
}

func DeleteUser(user_id string) error {
	err := DB.Model(models.User{}).Where("user_id = ? ", user_id).Update("isDelete", "1").Error
	if err != nil {
		log.Print("delete user error", err)
		return err
	}
	return nil
}

func SelectUserByUserName(username string) (*models.User, error) {
	var user models.User
	err := DB.Model(models.User{}).Where("username = ? and isDelete = 0", username).First(&user).Error
	if err != nil {
		log.Print("DB select user error", err)
		return nil, err
	}
	return &user, nil
}

func Insert(user models.User) error {
	return DB.Model(models.User{}).Create(&user).Error
}

func SelectUserByEmail(email string) int64 {
	var count int64
	err := DB.Model(models.User{}).
		Where("email = ?", email).
		Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return count
		}
		log.Print("select user by email count error : %v", err)
		return -1
	}
	return count
}
