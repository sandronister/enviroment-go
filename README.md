# enviroment-go

## Propósito do Projeto

O projeto `enviroment-go` tem como objetivo facilitar o carregamento e a gestão de variáveis de ambiente em aplicações Go. Ele fornece uma maneira simples e eficiente de carregar configurações a partir de arquivos `.env`, permitindo que as aplicações sejam configuradas de forma flexível e segura.

## Funcionalidades Principais

- Carregamento de variáveis ambiente a partir de arquivos `.env`.
- Estruturação de configurações em structs para fácil acesso e manipulação.
- Suporte para múltiplos serviços e configurações, como AWS e Redis.

## Exemplo de Uso

```go
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
        fmt.Printf("Server Name: %s\n", server.Name)
    }
}