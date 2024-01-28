package dao

import (
	"github.com/TJxiaobao/OJ/models"
	"log"
)

func GetUserDetail(user_id string) (*models.User, error) {
	var user models.User
	err := DB.Model(models.User{}).
		Where("user_id = ? and isDelete = 0", user_id).
		First(&user).Error
	if err != nil {
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
	err := DB.Model(models.User{}).First(&user).Where("user_name = ? and isDelete = 0", username).Error
	if err != nil {
		log.Print("DB select user error", err)
		return nil, err
	}
	return &user, nil
}

func Insert(user models.User) error {
	return DB.Model(models.User{}).Create(user).Error
}

func SelectUserByPhone(phone string) int64 {
	var count int64
	err := DB.Model(models.User{}).Count(&count)
	if err != nil {
		log.Print("select user by phone count error", err)
		return -1
	}
	return count
}
