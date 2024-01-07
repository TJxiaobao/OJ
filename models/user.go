package models

type User struct {
	BaseModel
	UserId   int    `gorm:"column:user_id" json:"-"`
	UserName string `gorm:"column:username" json:"-"`
	PassWord string `gorm:"column:password" json:"-"`
	Phone    int    `gorm:"column:phone" json:"-"`
	Email    string `gorm:"column:email" json:"-"`
	UserRole string `gorm:"column:user_role" json:"-"`
}

func (table User) TableName() string {
	return "oj_user"
}
