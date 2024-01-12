package cqe

import "github.com/TJxiaobao/OJ/utils/errno"

type GetUserDetailQuery struct {
	UserId string `json:"user_id"`
}

func (u *GetUserDetailQuery) Validate() error {
	if u.UserId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_id")
	}
	return nil
}

type DeleteUserCmd struct {
	UserId string `json:"user_id"`
}

func (u *DeleteUserCmd) Validate() error {
	if u.UserId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_id")
	}
	return nil
}

type LoginCmd struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	UserId   string `json:"user_id"`
}

func (l *LoginCmd) Validate() error {
	if l.UserId == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_id")
	}
	if l.UserName == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "user_name")
	}
	if l.PassWord == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "password")
	}
	return nil
}
