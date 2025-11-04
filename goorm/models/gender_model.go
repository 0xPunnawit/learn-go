package models

import "gorm.io/gorm"

type Gender struct {
	Id   uint
	Name string `gorm:"unique;size:10"`
}

type Test struct {
	gorm.Model
	Code uint   `gorm:"comment:This is a comment"`
	Name string `gorm:"column:myname;type:varchar(50);unique;default:'Hello';not null"`
}

func (t Test) TableName() string {
	return "MyTest"
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderId uint
}
