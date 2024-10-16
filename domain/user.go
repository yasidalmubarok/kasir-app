package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required"`
	Phone    string `form:"phone" json:"phone" binding:"required,phonenumber"`
	Email    string `form:"email" json:"email" binding:"required" gorm:"unique"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (u *User) HashPassword(*gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// func (u *User) Save() (*User, error) {
// 	var err error

// 	// check duplicate username
// 	err = Database.Database.Create(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}

// 	return u, nil
// }

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
