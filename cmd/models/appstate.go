package models

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppContext struct {
	echo.Context
	DB *gorm.DB
}