package po

import "gorm.io/gorm"

type StudentPo struct {
	gorm.Model
	StuId int    `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Major string `gorm:"column:major"`
}
