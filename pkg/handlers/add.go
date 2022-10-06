package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) addDriver(c *gin.Context) {
	var (
		param          string
		age            int8
		workExperience int8
		carId          int64
	)
	name := c.Query("name")
	secondName := c.Query("secondName")
	ageStr := c.Query("age")
	workExperienceStr := c.Query("workExperience")
	citizenship := c.Query("citizenship")
	address := c.Query("address")
	mobilePhone := c.Query("mobilePhone")
	email := c.Query("email")
	carIdStr := c.Query("carId")

	switch {
	case name == "":
		param = "name"
	case secondName == "":
		param = "secondName"
	case ageStr == "":
		param = "age"
	case workExperienceStr == "":
		param = "workExperience"
	case citizenship == "":
		param = "citizenship"
	case address == "":
		param = "address"
	case mobilePhone == "":
		param = "mobilePhone"
	case email == "":
		param = "email"
	case carIdStr == "":
		param = "carId"
	}

	if param != "" {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "missing parameter " + param})
		c.Abort()
		return
	}

	if vAge, err := strconv.ParseInt(ageStr, 10, 8); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "invalid parameter age"})
		c.Abort()
		return
	} else {
		age = int8(vAge)
	}

	if vWorkExperience, err := strconv.ParseInt(workExperienceStr, 10, 8); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "invalid parameter workExperience"})
		c.Abort()
		return
	} else {
		workExperience = int8(vWorkExperience)
	}

	if vCarId, err := strconv.ParseInt(carIdStr, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "invalid parameter carId"})
		c.Abort()
		return
	} else {
		carId = vCarId
	}

	err := h.Db.AddNewDriver(name, secondName, age, workExperience, citizenship,
		address, mobilePhone, email, carId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}

func (h *Handler) addCar(c *gin.Context) {
	var (
		param   string
		mileage int64
	)

	num := c.Query("num")
	model := c.Query("model")
	mileageStr := c.Query("mileage")

	switch {
	case num == "":
		param = "num"
	case model == "":
		param = "model"
	case mileageStr == "":
		param = "mileage"
	}

	if param == "" {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "missing parameter " + param})
		c.Abort()
		return
	}

	if vMileage, err := strconv.ParseInt(mileageStr, 10, 8); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"Error": "invalid parameter age"})
		c.Abort()
		return
	} else {
		mileage = vMileage
	}

	err := h.Db.AddNewCar(num, model, mileage)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
