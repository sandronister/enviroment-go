package main

import (
	"fmt"

	"github.com/sandronister/enviroment-go/pkg/load"
)

type AWS struct {
	Region string `json:"AWS_REGION"`
	Secret string `json:"AWS_SECRET_KEY_ID"`
	Key    string `json:"AWS_SECRET_ACCESS_KEY"`
}

type Redis struct {
	Port string `json:"REDIS_PORT"`
	Host string `json:"REDIS_HOST"`
	Db   int    `json:"REDIS_DATABASE"`
	Pass string `json:"REDIS_PASS"`
}

type Servers struct {
	Name   string `json:"Name"`
	ID     int    `json:"Id"`
	Port   int    `json:"Port"`
	Enable bool   `json:"Enable"`
}

func main() {
	varEnv := load.New("../.env")

	aws := AWS{}

	err := varEnv.LoadVariable(&aws)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", aws)

}
