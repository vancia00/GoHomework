package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm:"column:username" form:"username"  json:"username" binding:"required,max=10"`
	Email       string `gorm:"column:email" form:"email" json:"email" binding:"required"`
	Password    string `gorm:"column:password" form:"password" json:"password" binding:"required,min=8,max=20"`
	PhoneNumber string `gorm:"column:phoneNumber" form:"phoneNumber" json:"phoneNumber" binding:"required,len=11"`
	Salt        string `db:"salt" form:"salt" json:"salt"`
}

type UserRegister struct {
	UserName    string `gorm:"column:username" form:"username"  json:"username" binding:"required,max=10"`
	Email       string `gorm:"column:email" form:"email" json:"email" binding:"required"`
	Password    string `gorm:"column:password" form:"password" json:"password" binding:"required,min=8,max=20"`
	PhoneNumber string `gorm:"column:phoneNumber" form:"phoneNumber" json:"phoneNumber" binding:"required,len=11"`
}

type UserNameLogin struct {
	UserName string `gorm:"column:username" form:"username"  json:"username" binding:"required,max=10"`
	Password string `gorm:"column:password" form:"password" json:"password" binding:"required,min=8,max=20"`
}

type PhoneLogin struct {
	PhoneNumber string `gorm:"column:phoneNumber" form:"phoneNumber" json:"phoneNumber" binding:"required,len=10"`
	Password    string `gorm:"column:password" form:"password" json:"password" binding:"required,min=8,max=20"`
}

type EmailLogin struct {
	Email    string `gorm:"column:email" form:"email" json:"email" binding:"required"`
	Password string `gorm:"column:password" form:"password" json:"password" binding:"required,min=8,max=20"`
}

// 怎么获取到user里面的salt赋值到这个里面
type GetPassword struct {
	Password string `db:"password" form:"password" json:"password"`
	Salt     string `db:"salt" form:"salt" json:"salt"`
}
