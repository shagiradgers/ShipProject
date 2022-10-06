package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllDrivers(c *gin.Context) {
	drivers, err := h.Db.GetAllDrivers()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "res": drivers})
}

func (h *Handler) getAllCars(c *gin.Context) {
	cars, err := h.Db.GetAllCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "res": cars})

}

func (h *Handler) getDriver(c *gin.Context) {
	name := c.Query("name")
	secondName := c.Query("secondName")

	if name == "" {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "missing parameter name"})
		c.Abort()
		return
	}

	if secondName == "" {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "missing parameter secondName"})
		c.Abort()
		return
	}

	drivers, err := h.Db.GetDriverByNameAndSecondName(name, secondName)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "res": drivers})
}

func (h *Handler) getCar(c *gin.Context) {
	num := c.Query("num")

	if num == "" {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "missing parameter num"})
		c.Abort()
		return
	}

	cars, err := h.Db.GetCarsByNum(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "res": cars})

}
