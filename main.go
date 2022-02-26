package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed : %v", err)
	}

	viper.Set("redis.port", 8888)

	fmt.Println(viper.Get("app_name"))
	fmt.Println("mysql port:", viper.Get("mysql.port"))
	fmt.Println("mysql user:", viper.Get("mysql.user"))
	fmt.Println("mysql password:", viper.Get("mysql.password"))
	fmt.Println("mysql database:", viper.Get("mysql.database"))

	fmt.Println("redis ip:", viper.Get("redis.ip"))
	fmt.Println("redis port:", viper.Get("redis.port"))

	fmt.Println(viper.IsSet("redis.port"))
	fmt.Println(viper.IsSet("redis.port1"))
}
