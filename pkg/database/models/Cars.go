package models

import "fmt"

type Car struct {
	Id      int64  `json:"Id"`      // id
	Num     string `json:"Num"`     // номер машины
	Model   string `json:"Model"`   // модель машины
	Mileage int64  `json:"Mileage"` // пробег
}

// String имплементация интерфейса Stringer (fmt)
func (car *Car) String() string {
	return fmt.Sprintf("id: %v\n"+
		"Номер машины: %s\n"+
		"Модель машины: %s\n"+
		"Пробег: %v\n",
		car.Id, car.Num, car.Model, car.Mileage)
}
