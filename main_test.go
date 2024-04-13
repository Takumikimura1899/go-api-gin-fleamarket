package main

import (
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "Item1", Price: 1000, Description: "Item1 Description", SoldOut: false, UserID: 1},
		{Name: "Item2", Price: 2000, Description: "Item2 Description", SoldOut: true, UserID: 1},
		{Name: "Item3", Price: 3000, Description: "Item3 Description", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Item{})

	setupTestData(db)
	router := setUpRouter(db)

	return router
}
