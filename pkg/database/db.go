package database

import (
	"ShipProject/pkg/database/models"
	"database/sql"
	_ "modernc.org/sqlite"
)

const driverName = "sqlite"

type Database struct {
	filePath string // путь до файла бд
}

// NewDb создание нового экземпляра базы данных
func NewDb(filepath string) *Database {
	db := new(Database)
	db.filePath = filepath
	return db
}

// Init Проверка того, что база существует и работает
func (db *Database) Init() error {
	d, err := sql.Open(driverName, db.filePath)
	if err != nil {
		return err
	}

	defer func(d *sql.DB) {
		if err := d.Close(); err != nil {
			return
		}
	}(d)

	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

// GetAllDrivers Получить всех водителей
func (db *Database) GetAllDrivers() ([]models.Driver, error) {
	drivers := make([]models.Driver, 0)

	d, err := sql.Open(driverName, db.filePath)
	if err != nil {
		return nil, err
	}

	defer func(d *sql.DB) {
		if err := d.Close(); err != nil {
			return
		}
	}(d)

	rows, err := d.Query("select * from drivers")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		driver := models.Driver{}

		err := rows.Scan(&driver.Id, &driver.Name, &driver.SecondName, &driver.Age,
			&driver.WorkExperience, &driver.Citizenship, &driver.Address,
			&driver.MobilePhone, &driver.Email, &driver.Car)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	if err := d.Close(); err != nil {
		return nil, err
	}
	return drivers, nil
}

// GetAllCars Получить все машины
func (db *Database) GetAllCars() ([]models.Car, error) {
	cars := make([]models.Car, 0)

	d, err := sql.Open(driverName, db.filePath)
	if err != nil {
		return nil, err
	}

	defer func(d *sql.DB) {
		if err := d.Close(); err != nil {
			return
		}
	}(d)

	rows, err := d.Query("select * from cars")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		car := models.Car{}

		err := rows.Scan(&car.Id, &car.Num, &car.Model, &car.Mileage)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

// AddNewCar добавить новую машины в базу
func (db *Database) AddNewCar(num string, model string, mileage int64) error {
	d, err := sql.Open(driverName, db.filePath)
	if err != nil {
		return err
	}

	defer func(d *sql.DB) {
		if err := d.Close(); err != nil {
			return
		}
	}(d)

	_, err = d.Exec("insert into cars values(last_insert_rowid()+1, $1, $2, $3)",
		num, model, mileage)
	if err != nil {
		return err
	}

	return nil
}

// AddNewDriver добавить нового водителя в базу
func (db *Database) AddNewDriver(name string, secondName string, age int8,
	workExperience int8, citizenship string, address string, mobilePhone string,
	email string, carId int64) error {
	d, err := sql.Open(driverName, db.filePath)
	if err != nil {
		return err
	}

	defer func(d *sql.DB) {
		if err := d.Close(); err != nil {
			return
		}
	}(d)

	if carId == -1 {
		_, err = d.Exec("insert into drivers values(last_insert_rowid()+1, $1, $2, $3, $4, $5, $6, $7, $8, $9)",
			name, secondName, age, workExperience, citizenship, address, mobilePhone,
			email, nil)
	} else {
		_, err = d.Exec("insert into drivers values(last_insert_rowid()+1, $1, $2, $3, $4, $5, $6, $7, $8, $9)",
			name, secondName, age, workExperience, citizenship, address, mobilePhone,
			email, carId)
	}
	if err != nil {
		return err
	}

	return nil
}

// GetDriverByNameAndSecondName получение водителя по имения и фамилии
func (db *Database) GetDriverByNameAndSecondName(name, secondName string) ([]models.Driver, error) {
	d, err := sql.Open(driverName, db.filePath)
	if err != nil {
		return nil, err
	}

	defer func(d *sql.DB) {
		if err := d.Close(); err != nil {
			return
		}
	}(d)

	rows, err := d.Query("select * from drivers where name = $1 and second_name = $2",
		name, secondName)
	if err != nil {
		return nil, err
	}

	drivers := make([]models.Driver, 0)

	for rows.Next() {
		driver := models.Driver{}

		err := rows.Scan(&driver.Id, &driver.Name, &driver.SecondName, &driver.Age,
			&driver.WorkExperience, &driver.Citizenship, &driver.Address,
			&driver.MobilePhone, &driver.Email, &driver.Car)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	return drivers, nil
}

// GetCarsByNum получение машины по номеру
func (db *Database) GetCarsByNum(num string) ([]models.Car, error) {
	d, err := sql.Open(driverName, db.filePath)
	if err != nil {
		return nil, err
	}

	defer func(d *sql.DB) {
		if err := d.Close(); err != nil {
			return
		}
	}(d)

	rows, err := d.Query("select * from cars where num = $1", num)
	if err != nil {
		return nil, err
	}

	cars := make([]models.Car, 0)

	for rows.Next() {
		car := models.Car{}

		err := rows.Scan(&car.Id, &car.Num, &car.Model, &car.Mileage)
		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars, nil
}
