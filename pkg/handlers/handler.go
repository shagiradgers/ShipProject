package handlers

import (
	"ShipProject/pkg/database"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db *database.Database
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	get := router.Group("get") // группа получения данных
	{
		get.GET("all-drivers", h.getAllDrivers)
		get.GET("all-cars", h.getAllCars)

		get.GET("driver", h.getDriver)
		get.GET("car", h.getCar)
	}

	edit := router.Group("/add") // группа добавления данных
	{
		edit.POST("driver", h.addDriver)
		edit.POST("car", h.addCar)
	}

	return router
}
