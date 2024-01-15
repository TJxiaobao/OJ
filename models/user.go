package models

type User struct {
	BaseModel
	UserId   string `gorm:"column:user_id" json:"-"`
	UserName string `gorm:"column:username" json:"-"`
	PassWord string `gorm:"column:password" json:"-"`
	Phone    string `gorm:"column:phone" json:"-"` // 字段需要修改
	Email    string `gorm:"column:email" json:"-"`
	UserRole int    `gorm:"column:user_role" json:"-"` // 数据库数值默认都为0， 即普通用户
}

func (table User) TableName() string {
	return "oj_user"
}
