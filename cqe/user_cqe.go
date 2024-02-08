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

type LoginGitHubQuery struct {
	Code string `json:"code" form:"code"`
}

func (l *LoginGitHubQuery) Validate() error {
	if l.Code == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code")
	}
	return nil
}

type Register struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
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
	if r.Email == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "email")
	}
	if r.Code == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code")
	}
	if r.Email == "" && r.Phone == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "email or phone must not null")
	}
	return nil
}

type SendCodeCmdByEmail struct {
	Email string `json:"email" form:"email"`
}

func (s *SendCodeCmdByEmail) Validate() error {
	if s.Email == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "email")
	}
	return nil
}

type SendCodeCmdBySms struct {
	Phone string `json:"phone" form:"phone"`
}

func (s *SendCodeCmdBySms) Validate() error {
	if s.Phone == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "phone")
	}
	return nil
}

type LoginSmsCmd struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}

func (l *LoginSmsCmd) Validate() error {
	if l.Phone == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "phone")
	}
	if l.Code == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code")
	}
	return nil
}
