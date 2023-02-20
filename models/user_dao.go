package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func GetUserById(id int64) (*User, error) {
	var user User
	result := db.Model(&User{}).Where("id=?", id).Find(&user)
	if result.Error != nil {
		return &User{}, result.Error
	}
	return &user, nil
}

func IsExistByName(name string) (bool, error) {
	var count int64
	result := db.Model(&User{}).Where("`name`=?", name).Count(&count)
	if result.Error != nil {
		return true, result.Error
	}
	return count == 1, nil
}

func CreateUser(name, password string) (*User, error) {
	user := User{Name: name, Password: password}
	result := db.Model(&User{}).Create(&user)
	if result.Error != nil {
		return &User{}, result.Error
	}
	return &user, nil
}

func IsExistByNameAndPassword(name, password string) (bool, int64, error) {
	var user User
	result := db.Model(&User{}).Where("`name`=? and `password`=?", name, password).Find(&user)
	if result.Error != nil {
		return true, 0, result.Error
	}
	if user.ID != 0 {
		return true, int64(user.ID), nil
	}
	return false, 0, nil
}
