package users

import "gorm.io/gorm"

type users struct {
	gorm.Model
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserInfo struct {
	gorm.Model
	Address_1 string 
	Address_2 string
	City string
	CountryCode string
	Email string
	FirstName string
	LastName string
	Phone string
	PostalCode string
	State string
	AddressID int
}