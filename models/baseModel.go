package models

import "time"

type BaseModel struct {
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:createTime" json:"created_time"`
	UpdatedAt time.Time `gorm:"column:updateTime" json:"updated_time"`
	IsDeleted int64     `gorm:"column:isDelete" json:"-"`
}
