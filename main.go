package main

import (
	"fmt"

	"github.com/dixneuf19/go-min-docker/config"
)

func main() {
	config.Init()
	fmt.Println("App initiated")
}
