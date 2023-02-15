package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func GetUserById(id int64) (User, error) {
	var user User
	result := db.Model(&User{}).Where("id=?", id).Find(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
