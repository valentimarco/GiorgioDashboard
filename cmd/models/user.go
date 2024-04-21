package models


type User struct {
	username string `form:"username"`
	password string `form:"password"`
}