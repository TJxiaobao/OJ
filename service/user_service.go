package service

import (
	"github.com/TJxiaobao/OJ/cqe"
	"github.com/TJxiaobao/OJ/dao"
	"github.com/TJxiaobao/OJ/utils/md5"
	"github.com/TJxiaobao/OJ/utils/restapi"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserDetail(c *gin.Context) {
	query := cqe.GetUserDetailQuery{}
	if err := c.ShouldBindJSON(query); err != nil {
		log.Print("user query error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 user_id 是否为空
	if err := query.Validate(); err != nil {
		log.Print("user_id must not null", err)
		restapi.Failed(c, err)
		return
	}

	user, err := dao.GetUserDetail(query.UserId)
	if err != nil {
		restapi.FailedWithStatus(c, err, http.StatusInternalServerError)
		return
	}
	if user == nil {
		restapi.Success(c, "the user does exist, please register !")
		return
	}
	restapi.Success(c, user)
	return
}

func DeleteUser(c *gin.Context) {
	cmd := cqe.DeleteUserCmd{}
	if err := c.ShouldBindJSON(cmd); err != nil {
		log.Print("user query error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 user_id 是否为空
	if err := cmd.Validate(); err != nil {
		log.Print("user_id must not null", err)
		restapi.Failed(c, err)
		return
	}

	err := dao.DeleteUser(cmd.UserId)
	if err != nil {
		restapi.FailedWithStatus(c, err, http.StatusInternalServerError)
		return
	}
	restapi.Success(c, "删除成果！")
}

func Login(c *gin.Context) {
	cmd := cqe.LoginCmd{}
	if err := c.ShouldBindJSON(cmd); err != nil {
		log.Print("Login cmd error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 参数 是否为空
	if err := cmd.Validate(); err != nil {
		log.Print("Login params must not null", err)
		restapi.Failed(c, err)
		return
	}

	data, err := dao.SelectUserByUserName(cmd.UserName)
	if err != nil {
		restapi.FailedWithStatus(c, err, 500)
		return
	}
	password := md5.Md5Encrypt(cmd.PassWord)
	if data.PassWord != password {
		restapi.Success(c, "password error")
		return
	} else {
		token := restapi.NewTokenResult("token")
		restapi.Success(c, token)
		return
	}
}
