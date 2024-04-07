package main

import (
	"gin-fleamarket/controller"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "Item 1", Price: 100, Description: "This is item 1", SoldOut: false},
		{ID: 2, Name: "Item 2", Price: 200, Description: "This is item 2", SoldOut: true},
		{ID: 3, Name: "Item 3", Price: 300, Description: "This is item 3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controller.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run()
}
