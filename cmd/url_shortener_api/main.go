package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"url_shortener_api/internal/app/url_shortener_api/server"
)

func init() {
	viper.SetConfigFile("configs/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("host=%s dbname=%s sslmode=disable port=%s user=%s password=%s ", dbHost, dbName, dbPort, dbUser, dbPass)

	fmt.Println("Server running")

	config := server.Config{
		DatabaseURL:  connection,
		Address:      viper.GetString("server.address"),
		Timeout:      viper.GetInt("context.timeout"),
		Ttl:          viper.GetInt("ttl"),
		CronInterval: viper.GetInt("cron_interval"),
	}

	err := server.StartApp(config)
	if err != nil {
		log.Fatal(err)
	}
}
