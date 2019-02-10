package util

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	return
}
