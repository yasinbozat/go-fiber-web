package models

import (
	"go-web/database"
	"strconv"
)

type User struct {
	Id          int    `gorm:"primary_key;auto_increment"`
	Name        string `gorm:"size:50"`
	Surname     string `gorm:"size:50"`
	Mail        string `gorm:"size:100;not_null;unique"`
	Password    string `gorm:"size:50;not_null"`
	PhoneNumber string `gorm:"size:25"`
	Country     string `gorm:"size:50"`
}

func SelectUser(limit, offset int) (myString string) {
	var tbuser []User
	database.DB.Limit(limit).Offset(offset).Find(&tbuser)
	for _, user := range tbuser {

		myString += "{Id: '" + strconv.Itoa(user.Id) + "', Name: '" + user.Name + "', Mail: '" + user.Mail +
			"', PhoneNumber: '" + user.PhoneNumber + "', Country: '" + user.Country + "'},"

	}
	myString = "[" + myString + "]"
	return myString
}