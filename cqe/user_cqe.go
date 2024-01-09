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
