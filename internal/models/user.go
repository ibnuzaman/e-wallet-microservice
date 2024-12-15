package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username" gorm:"column:username;type:varchar(20)" validate:"required"`
	Email       string `json:"email" gorm:"column:email;type:varchar(100)" validate:"required,email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number;type:varchar(15)" validate:"required"`
	FullName    string `json:"full_name" gorm:"column:full_name;type:varchar(100)"`
	Address     string `json:"address" gorm:"column:address;type:text"`
	Dob         string `json:"dob" gorm:"column:dob;type:date"`
	Password    string `json:"password" gorm:"column:password;type:varchar(100)" validate:"required"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
}

func (*User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  uint `gorm:"primary_key"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              uint   `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string `json:"token" gorm:"type:varchar(100)" validate:"required"`
	RefreshToken        string `json:"refresh_token" gorm:"type:varchar(100)" validate:"required"`
	TokenExpired        string `json:"-" validate:"required"`
	RefreshTokenExpired string `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
