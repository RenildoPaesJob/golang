package main

import (
	"fmt"

	"github.com/RenildoPaesJob/golang.git/config"
	"github.com/RenildoPaesJob/golang.git/router"
)

func main() {
	// print("A galera da live é braba e vai conseguir trampo de 200k/mês!")

	//initialize Config
	err := config.Init()

	if err != nil {
		fmt.Println(err)
	}

	//Initialize Router
	router.Inicialize()
}