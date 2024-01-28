package cqe

import "github.com/TJxiaobao/OJ/utils/errno"

type GetUserDetailQuery struct {
	UserId string `json:"user_id" form:"user_id"`
}

func (u *GetUserDetailQuery) Validate() error {
	if u.UserId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_id")
	}
	return nil
}

type DeleteUserCmd struct {
	UserId string `json:"user_id" form:"user_id"`
}

func (u *DeleteUserCmd) Validate() error {
	if u.UserId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_id")
	}
	return nil
}

type LoginCmd struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func (l *LoginCmd) Validate() error {
	if l.UserName == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_name")
	}
	if l.PassWord == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "password")
	}
	return nil
}

type Register struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Code     string `json:"code" form:"phone"`
}

func (r *Register) Validate() error {
	if r.UserName == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "username")
	}
	if r.Password == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "passord")
	}
	if r.Phone == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "phone")
	}
	if r.Code == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code")
	}
	return nil
}

type SendCodeCmd struct {
	Phone string `json:"phone" form:"phone"`
}

func (s *SendCodeCmd) Validate() error {
	if s.Phone == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "phone")
	}
	return nil
}
