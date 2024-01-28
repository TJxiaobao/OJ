package models

type User struct {
	BaseModel
	UserId   string `gorm:"column:user_id" json:"user_id"`
	UserName string `gorm:"column:username" json:"user_name"`
	PassWord string `gorm:"column:password" json:"pass_word"`
	Phone    string `gorm:"column:phone" json:"phone"` // 字段需要修改
	Email    string `gorm:"column:email" json:"email"`
	UserRole string `gorm:"column:user_role" json:"user_role"` // 数据库数值默认都为0， 即普通用户
}

func (table User) TableName() string {
	return "oj_user"
}
