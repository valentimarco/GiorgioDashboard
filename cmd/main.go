package main

import (
	"os"
	"website/cmd/models"
	"website/cmd/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Debug("Execution on container/pod??? Check your execution!!")
	}
	url_db := os.Getenv("POSTGRES_HOST")
	sqlDB, err := sql.Open("pgx", url_db)
	if err != nil {
		panic(err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//migrations
	gormDB.AutoMigrate(&models.User{})

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &models.AppContext{
				Context: c,
				DB: gormDB,
			}
			return next(cc)
		}
	})
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.Static("/css", "css")
	e.Static("/public", "public")
	e.Static("/static", "static")
	routes.IndexRoutes(e)
	routes.BackendRoutes(e, gormDB)

	e.Logger.Fatal(e.Start(":3000"))
}
