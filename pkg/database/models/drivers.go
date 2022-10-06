package models

import "fmt"

type Driver struct {
	Id             int64  `json:"Id"`             // id
	Name           string `json:"Name"`           // Имя
	SecondName     string `json:"SecondName"`     // Фамилия
	Age            int8   `json:"Age"`            // Возраст
	WorkExperience int8   `json:"WorkExperience"` // Опыт работы
	Citizenship    string `json:"Citizenship"`    // гражданство
	Address        string `json:"Address"`        // адрес проживания
	MobilePhone    string `json:"MobilePhone"`    // мобильный телефон
	Email          string `json:"Email"`          // Почта
	Car            *Car   `json:"Car"`            // Машина
}

// String имплементация интерфейса Stringer (fmt)
func (dr *Driver) String() string {
	return fmt.Sprintf("id: %v\n"+
		"Имя: %s\n"+
		"Фамилия: %s\n"+
		"Возраст: %v\n"+
		"Опыт работы: %v\n"+
		"Гражданство: %s\n"+
		"Адрес: %s\n"+
		"Номер телефона: %s\n"+
		"Почта: %s\n",
		dr.Id, dr.Name, dr.SecondName,
		dr.Age, dr.WorkExperience,
		dr.Citizenship, dr.Address,
		dr.MobilePhone, dr.Email)
}
