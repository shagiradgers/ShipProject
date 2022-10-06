package main

import (
	"ShipProject/pkg/database"
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

func printAllDrivers(db *database.Database, scanner *bufio.Scanner) {
	drivers, err := db.GetAllDrivers()
	if err != nil {
		log.Fatal("Error, then getting drivers: ", err)
	}

	if len(drivers) == 0 {
		fmt.Println("Пока водителей нет")
	}

	for _, v := range drivers {
		fmt.Println(&v)
	}
	fmt.Println("Продолжить? 1-да, 0-нет")
	scanner.Scan()
	if scanner.Text() != "1" {
		os.Exit(1)
	}
}

func printAllCars(db *database.Database, scanner *bufio.Scanner) {
	cars, err := db.GetAllCars()
	if err != nil {
		log.Fatal("Error, then getting cars: ", err)
	}

	if len(cars) == 0 {
		fmt.Println("Пока машин нет")
	}

	for _, v := range cars {
		fmt.Println(&v)
	}

	fmt.Println("Продолжить? 1-да, 0-нет")
	scanner.Scan()
	if scanner.Text() != "1" {
		os.Exit(1)
	}
}

func findDriver(db *database.Database, scanner *bufio.Scanner) {
	fmt.Println("Отправьте имя водителя")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Отправьте фамлию водителя")
	scanner.Scan()
	secondName := scanner.Text()

	drivers, err := db.GetDriverByNameAndSecondName(name, secondName)
	if err != nil {
		log.Fatal("Error, then getting drivers by parameters: ", err)
	}

	if len(drivers) == 0 {
		fmt.Println("Ничего не нашли")
	}

	for _, v := range drivers {
		fmt.Print(&v, "\n")
	}

	fmt.Println("Продолжить? 1-да, 0-нет")
	scanner.Scan()
	if scanner.Text() != "1" {
		os.Exit(1)
	}

}

func findCar(db *database.Database, scanner *bufio.Scanner) {
	fmt.Println("Отправьте номер машины")
	scanner.Scan()
	num := scanner.Text()

	cars, err := db.GetCarsByNum(num)

	if err != nil {
		fmt.Println("Error, then getting cars by parameters", err)
	}

	if len(cars) == 0 {
		fmt.Println("Ничего не нашли")
	}

	for _, v := range cars {
		fmt.Print(&v, "\n")
	}

	fmt.Println("Продолжить? 1-да, 0-нет")
	scanner.Scan()
	if scanner.Text() != "1" {
		os.Exit(1)
	}
}

func addDriver(db *database.Database, scanner *bufio.Scanner) {
	fmt.Println("Отправьте имя водителя")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Отправьте фамилию водителя")
	scanner.Scan()
	secondName := scanner.Text()

	fmt.Println("отправьте возраст водителя")
	scanner.Scan()
	age, err := strconv.ParseInt(scanner.Text(), 10, 8)
	if err != nil {
		fmt.Println("Вы ввели не число, я обиделся")
		return
	}

	fmt.Println("Отправьте опыт работы водителя")
	scanner.Scan()
	workExperience, err := strconv.ParseInt(scanner.Text(), 10, 8)
	if err != nil {
		fmt.Println("Вы ввели не число, я обиделся")
		return
	}

	fmt.Println("Отправьте гражданство водителя")
	scanner.Scan()
	citizenship := scanner.Text()

	fmt.Println("Отправьте адрес проживания водителя")
	scanner.Scan()
	address := scanner.Text()

	fmt.Println("Отправьте номер телефона водителя")
	scanner.Scan()
	mobilePhone := scanner.Text()

	fmt.Println("Отправьте почту водителя")
	scanner.Scan()
	email := scanner.Text()

	fmt.Println("Отправьте id машины водителя (если машины нет в базе, " +
		"то отправьте -1)")
	scanner.Scan()
	carId, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println("Вы ввели не число, я обиделся")
		return
	}

	err = db.AddNewDriver(name, secondName, int8(age), int8(workExperience),
		citizenship, address, mobilePhone, email, carId)
	if err != nil {
		log.Fatal("Error, then add new driver: ", err)
	}
}

func addCar(db *database.Database, scanner *bufio.Scanner) {
	fmt.Println("Отправьте номер машины")
	scanner.Scan()
	num := scanner.Text()

	fmt.Println("Отправьте модель машины")
	scanner.Scan()
	model := scanner.Text()

	fmt.Println("Отправьте пробег машины (в км)")
	scanner.Scan()
	mileage, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println("Вы ввели не число, я обиделся")
		return
	}

	if err := db.AddNewCar(num, model, mileage); err != nil {
		log.Fatal("Error, then add new car: ", err)
	}
}

func main() {
	fmt.Println("Запускам волшебную штуку")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Читаем конфиги...")
	if err := initConfig(); err != nil {
		log.Fatal("Error, then reading config: ", err.Error())
	}

	db := database.NewDb(viper.GetString("dbFilePath"))

	fmt.Println("Проверяем базу...")
	if err := db.Init(); err != nil {
		log.Fatal("Error, then initializing db: ", err.Error())
	}

	for {
		fmt.Print("\x1b[2J") // очистка консоли

		fmt.Println("1. Вывести всех водителей")
		fmt.Println("2. Вывести все машины")
		fmt.Println("3. Найти водителя")
		fmt.Println("4. Найти машину")
		fmt.Println("5. Добавить водителя")
		fmt.Println("6. Добавить машину")
		fmt.Println("0. Выйти")

		scanner.Scan()

		switch scanner.Text() {
		case "1":
			printAllDrivers(db, scanner)
		case "2":
			printAllCars(db, scanner)
		case "3":
			findDriver(db, scanner)
		case "4":
			findCar(db, scanner)
		case "5":
			addDriver(db, scanner)
		case "6":
			addCar(db, scanner)
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Чего?")
		}
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
