package conf

import (
	"os"
	"strconv"
)

type Config struct {
	Port uint16
}

func NewConfig() Config {
	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)

	if err != nil {
		panic(err)
	}
	return Config{Port: uint16(port)}
}
