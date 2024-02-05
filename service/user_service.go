package service

import (
	"fmt"
	"github.com/TJxiaobao/OJ/constant"
	"github.com/TJxiaobao/OJ/cqe"
	"github.com/TJxiaobao/OJ/dao"
	"github.com/TJxiaobao/OJ/models"
	"github.com/TJxiaobao/OJ/utils/errno"
	"github.com/TJxiaobao/OJ/utils/md5"
	"github.com/TJxiaobao/OJ/utils/restapi"
	"github.com/TJxiaobao/OJ/utils/token"
	"github.com/TJxiaobao/OJ/utils/uuid"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func GetUserDetail(c *gin.Context) {
	query := cqe.GetUserDetailQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
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
	if err := c.BindJSON(&cmd); err != nil {
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
	restapi.Success(c, "删除成功！")
}

// Login
// @Summary      用户登陆
// @Description	 test
// @Tags         user
// @Param        user_id   formData string  false "user_id"
// @Param        username   formData string  false "username"
// @Param        password   formData string  false "password"
// @Success      200  {string} json "{"code" : "200", "msg" : "", "data" : ""}"
// @Router       /user/login [post]
func Login(c *gin.Context) {
	cmd := cqe.LoginCmd{}
	if err := c.BindJSON(&cmd); err != nil {
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
		token_str, err := token.GenerateToken(data.UserName, data.UserId)
		if err != nil {
			log.Print("generate token error", err)
			restapi.FailedWithStatus(c, err, 500)
			return
		}
		token_result := restapi.NewTokenResult(token_str)
		restapi.Success(c, token_result)
		return
	}
}

func Register(c *gin.Context) {
	cmd := cqe.Register{}
	if err := c.BindJSON(&cmd); err != nil {
		log.Print("register cmd error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 参数 是否为空
	if err := cmd.Validate(); err != nil {
		log.Print("Register params must not null", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 验证码 是否正确
	// todo test
	deCode, err := dao.InitRedis().Get(c, cmd.Email).Result()
	if err != nil {
		restapi.FailedWithStatus(c, err, 500)
		return
	}
	if deCode != cmd.Code {
		restapi.FailedWithStatus(c, errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code"), 400)
		return
	}

	// 通过手机号判断是否已经注册
	count := dao.SelectUserByEmail(cmd.Email)
	if count > 0 {
		restapi.Success(c, "该邮箱已注册！")
		return
	}

	// 注册用户
	user := models.User{
		UserId:   uuid.GetUUID(),
		UserName: cmd.UserName,
		PassWord: md5.Md5Encrypt(cmd.Password),
		Email:    cmd.Email,
		UserRole: constant.UserNormalRole,
	}
	err = dao.Insert(user)
	if err != nil {
		log.Print("db insert error : ", err)
		restapi.FailedWithStatus(c, err, 500)
		return
	}

	// 生成 token
	newToken, err := token.GenerateToken(user.UserName, user.UserId)
	resp := restapi.NewTokenResult(newToken)
	restapi.Success(c, resp)
}

// createCode 生成6位验证码
func createCode() (code string) {
	rand.Seed(time.Now().Unix()) //设置随机种子
	code = fmt.Sprintf("%6v", rand.Intn(600000))
	return
}

func SendCode(c *gin.Context) {
	cmd := cqe.SendCodeCmd{}
	if err := c.BindJSON(&cmd); err != nil {
		log.Print("sendCode cmd error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 参数 是否为空
	if err := cmd.Validate(); err != nil {
		log.Print("sendcode params must not null", err)
		restapi.Failed(c, err)
		return
	}

	code := createCode()
	err := dao.InitRedis().Set(c, cmd.Email, code, time.Second*300) // 设置时间为 5 分钟
	if err != nil {
		log.Print("redis set error", err)
		restapi.FailedWithStatus(c, nil, 500)
		return
	}
	restapi.Success(c, "send code success !")
}
