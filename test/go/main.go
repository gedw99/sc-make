package main

import (
	zerolog "github.com/rs/zerolog/log"

	"github.com/joho/godotenv"

	"github.com/gedw99/gio-make/test/pkg/env"
)

func main() {
	zerolog.Printf("hello")

	err := godotenv.Load()
	if err != nil {
		zerolog.Print("Error loading .env file")
	}

	PORT := env.GetEnv("PORT", "4080")
	zerolog.Printf("PORT %v", PORT)

}
