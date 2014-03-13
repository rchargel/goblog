package web

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	port   uint16
	router *Router
}

func NewConfig() *Config {
	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)

	if err != nil {
		panic(err)
	}

	ch := make(chan bool)
	completed := 0

	router := NewRouter()
	go router.Init(ch)

	for success := range ch {
		completed += 1
		if !success {
			fmt.Println("Configuration Failed")
		}
		if completed == 1 {
			break
		}
	}

	return &Config{port: uint16(port), router: router}
}
