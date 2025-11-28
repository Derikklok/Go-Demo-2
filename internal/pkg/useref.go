package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func UseRef() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	PORT := os.Getenv("PORT")
	log.Println(PORT)
}
