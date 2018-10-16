package main

import (
	"fmt"
	"time"

	"github.com/dixneuf19/go-min-docker/config"
)

func main() {
	config.Init()
	fmt.Println("App initiated")

	for {
		time.Sleep(60 * time.Second)
	}
}
