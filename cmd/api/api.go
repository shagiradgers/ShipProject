package main

import (
	"ShipProject/pkg/database"
	"ShipProject/pkg/handlers"
	"ShipProject/pkg/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Error, then reading config ", err.Error())
	}
	db := database.NewDb(viper.GetString("dbFilePath")) // база данных

	if err := db.Init(); err != nil {
		log.Fatal("Error, then initializing db ", err.Error())
	}

	serv := new(server.Server)    // сервер
	hand := new(handlers.Handler) // обработчик запросов
	hand.Db = db                  // ссылка на базу данных в обработчике

	if err := serv.Run(viper.GetString("port"), hand.InitRoutes()); err != nil {
		log.Fatal("Some error occurred: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
