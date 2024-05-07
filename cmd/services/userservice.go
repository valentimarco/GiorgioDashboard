package services

import (
	"website/cmd/models"

	"gorm.io/gorm"
)

func InsertUser(userdto *models.UserDTO, db *gorm.DB) (statusCode int){
	user := models.User{
		Username: userdto.Username,
		Password: userdto.Password,
	}

	result := db.Create(&user)

	if result.Error != nil{
		statusCode = 500
	}

	statusCode = 200
	return
}

func FindUser(userdto *models.UserDTO, db *gorm.DB) (statusCode int){
	user := models.User{}

	result := db.Where(&models.User{Username: userdto.Username, Password: userdto.Password}).Find(&user)

	//Idk if there are gonna be more err...
	switch result.Error{
		case gorm.ErrRecordNotFound:
			statusCode = 404
		default:
			statusCode = 500
	}

	statusCode = 200
	return
}