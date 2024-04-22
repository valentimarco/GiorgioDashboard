package main

import (
	"log"
	"os"
	"website/src/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	e := echo.New()
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
