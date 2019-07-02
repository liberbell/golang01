package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("error: can`t open config file - %s\n", err)
	}

	defer file.Close()

	cfg := &Config{}
	if err := toml.NewDecoder(file); err != nil {
		log.Fatalf("error: can`t decode configuration file - %s", err)
	}
	fmt.Printf("%+v\n", cfg)
}
