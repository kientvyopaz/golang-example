package models

import (
	"golang-example/config"
	"strconv"
)

type User struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Age     uint   `gorm:"not null"`
	Phone   string `gorm:"not null"`
	Email   string `gorm:"not null"`
	Address string `gorm:"not null"`
}

func CreateUser(user *User) error {
	err := config.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User) error {
	err := config.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User) error {
	err := config.DB.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(id uint) (*User, error) {
	var user User
	err := config.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func SearchUsers(name string, ageStr string, phone string, email string, idStr string) ([]User, error) {
	age, _ := strconv.Atoi(ageStr)
	id, _ := strconv.Atoi(idStr)

	var users []User
	query := config.DB.Where("1=1")
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if age > 0 {
		query = query.Where("age = ?", age)
	}
	if phone != "" {
		query = query.Where("phone = ?", phone)
	}
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if id > 0 {
		query = query.Where("id = ?", id)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
