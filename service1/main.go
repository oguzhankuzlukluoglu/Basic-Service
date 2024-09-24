package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/oguzhankuzlukluoglu/service1/config"
)

func main() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// DATABASE_URL değişkenini al
	databaseURL := os.Getenv("DATABASE_URL")

	// Veritabanı bağlantısını kur
	db, err := config.ConnectToDB(databaseURL)
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız: %v", err)
	}
	defer db.Close()
	fmt.Println("Veritabanına başarıyla bağlandım!")
}
