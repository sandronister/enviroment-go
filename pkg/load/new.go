package load

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sandronister/enviroment-go/pkg/types"
)

type environment struct {
	variables map[string]string
}

func loadMap(path string) (map[string]string,error) {
	err := godotenv.Load(path)

	if err != nil {
		return nil, err
	}

	envVars := os.Environ()
	variables := make(map[string]string)

	for _, value := range envVars {
		tmp := strings.Split(value, "=")
		variables[tmp[0]] = tmp[1]
	}

	return variables, nil
}

func New(path string) (types.EnvironmentPorts,error) {
	variable,err := loadMap(path)
	if err != nil {
		return nil, err
	}
	
	return &environment{
		variables: variable,
	},nil
}
