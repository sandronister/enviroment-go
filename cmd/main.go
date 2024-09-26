package main

import (
	"fmt"

	"github.com/sandronister/enviroment-go/pkg/load"
)

type AWS struct {
	Region string `var:"AWS_REGION"`
	Secret string `var:"AWS_SECRET_KEY_ID"`
	Key    string `var:"AWS_SECRET_ACCESS_KEY"`
}

type Redis struct {
	Port string `var:"REDIS_PORT"`
	Host string `var:"REDIS_HOST"`
	Db   int    `var:"REDIS_DATABASE"`
	Pass string `var:"REDIS_PASS"`
}

type Servers struct {
	Name   string `json:"SERVER_NAME"`
	ID     int    `json:"SERVER_ID"`
	Port   int    `json:"SERVER_PORT"`
	Enable bool   `json:"SERVER_ENABLE"`
}

func main() {
	varEnv := load.New("../.env")

	var (
		aws     = AWS{}
		redis   = Redis{}
		servers = []Servers{}
	)

	errs := varEnv.Load(&aws, &redis)

	if len(errs) > 0 {
		for _, item := range errs {
			fmt.Printf("Error %s\n", item)
		}
	}

	fmt.Printf("%+v\n", aws)
	fmt.Printf("%v\n", redis)

	err := varEnv.LoadList("SERVERS", &servers)

	if err != nil {
		fmt.Println("Erro", err)
	}

	for _, server := range servers {
		fmt.Printf("%v\n", server)
	}

}
