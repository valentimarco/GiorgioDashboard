package models

import "gorm.io/gorm"


type UserDTO struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	AdminLevel int 

}
