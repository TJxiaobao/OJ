package service

import (
	"fmt"
	"github.com/TJxiaobao/OJ/constant"
	"github.com/TJxiaobao/OJ/cqe"
	"github.com/TJxiaobao/OJ/dao"
	"github.com/TJxiaobao/OJ/models"
	"github.com/TJxiaobao/OJ/utils/errno"
	"github.com/TJxiaobao/OJ/utils/md5"
	"github.com/TJxiaobao/OJ/utils/oauth/github"
	"github.com/TJxiaobao/OJ/utils/restapi"
	"github.com/TJxiaobao/OJ/utils/token"
	"github.com/TJxiaobao/OJ/utils/uuid"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
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

func LoginGetGithubUrl(c *gin.Context) {
	url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", constant.GithubOauthClientId, constant.GithubOauthRedirectUrl)
	c.JSON(200, url)
}

func LoginGitHub(c *gin.Context) {
	query := cqe.LoginGitHubQuery{}
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Print("get code error: ", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 参数 是否为空
	if err := query.Validate(); err != nil {
		log.Print("code must not null", err)
		restapi.Failed(c, err)
		return
	}

	// 获取url
	url := github.GetTokenAuthUrl(query.Code)

	// 获取token
	token, err := github.GetToken(url)
	if err != nil {
		log.Print("get token error: ", err)
		restapi.Failed(c, err)
		return
	}

	// 获取用户信息
	GithubUser, err := github.GetUserInfo(token)
	if err != nil {
		log.Print("get github user info error: ", err)
		restapi.Failed(c, err)
		return
	}

	print(GithubUser)
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
	var code string
	if cmd.Email == "" {
		code = cmd.Phone
	} else {
		code = cmd.Email
	}
	deCode, err := dao.InitRedis().Get(c, code).Result()
	if err != nil {
		restapi.FailedWithStatus(c, err, 500)
		return
	}
	if deCode != cmd.Code {
		restapi.FailedWithStatus(c, errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code"), 400)
		return
	}

	// 通过邮箱判断是否已经注册
	if cmd.Email != "" {
		count := dao.SelectUserByEmail(cmd.Email)
		if count > 0 {
			restapi.Success(c, "该邮箱已注册！")
			return
		}
	}

	// 判断手机号是否注册
	if cmd.Phone != "" {
		count := dao.SelectUserByPhone(cmd.Phone)
		if count > 0 {
			restapi.Success(c, "该手机号已经注册！")
			return
		}
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

func SendCodeByEmail(c *gin.Context) {
	cmd := cqe.SendCodeCmdByEmail{}
	if err := c.BindJSON(&cmd); err != nil {
		log.Print("sendCode Email cmd error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 参数 是否为空
	if err := cmd.Validate(); err != nil {
		log.Print("send_code_email params must not null", err)
		restapi.Failed(c, err)
		return
	}

	code := createCode()
	err := dao.InitRedis().Set(c, cmd.Email, code, time.Second*300).Err() // 设置时间为 5 分钟
	if err != nil {
		log.Print("redis set error", err)
		restapi.FailedWithStatus(c, nil, 500)
		return
	}

	//接收者邮箱列表
	mailTo := []string{
		cmd.Email,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", constant.EmailSender)          //发送者腾讯邮箱账号
	m.SetHeader("To", mailTo...)                       //接收者邮箱列表
	m.SetHeader("Subject", constant.EmailTitl)         //邮件标题
	m.SetBody("text/html", constant.EmailBodyTem+code) //邮件内容,可以是html

	//发送邮件服务器、端口、发送者qq邮箱、qq邮箱授权码
	//服务器地址和端口是腾讯的
	d := gomail.NewDialer(constant.EmailQQHost, 587, constant.EmailSender, constant.EmailAuthCode)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	restapi.Success(c, "send code success !")
}

func SendCodeBySms(c *gin.Context) {
	cmd := cqe.SendCodeCmdBySms{}
	if err := c.BindJSON(&cmd); err != nil {
		log.Print("sendCode Email cmd error", err)
		restapi.Failed(c, err)
		return
	}

	// 判断 参数 是否为空
	if err := cmd.Validate(); err != nil {
		log.Print("send_code_email params must not null", err)
		restapi.Failed(c, err)
		return
	}

	// 手机号校验

	// 存储到redis
	code := createCode()
	err := dao.InitRedis().Set(c, cmd.Phone, code, time.Second*300).Err() // 设置时间为 5 分钟
	if err != nil {
		log.Print("redis set error", err)
		restapi.FailedWithStatus(c, nil, 500)
		return
	}

	// 发送验证码
	// todo
}
