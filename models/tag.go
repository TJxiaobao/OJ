package models

type Tag struct {
	BaseModel
	ParentId int    `gorm:"column:parent_id" json:"-"`
	TagName  string `gorm:"column:tag_name" json:"-"`
}

func (table Tag) TableName() string {
	return "oj_type"
}
