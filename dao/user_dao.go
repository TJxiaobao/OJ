package dao

import (
	"github.com/TJxiaobao/OJ/models"
	"log"
)

func GetUserDetail(user_id string) (*models.User, error) {
	var user models.User
	err := DB.Model(models.User{}).
		First(&user).
		Where("user_id = ? and isDelete == 0", user_id).Error
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
