package db

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	database_url := os.Getenv("DATABASE_URL")
	regex := regexp.MustCompile("^postgres://(.+):(.+)@(.+):(.+)/(.+)$")
	result := regex.FindSubmatch([]byte(database_url))

	username := string(result[1])
	password := string(result[2])
	host := string(result[3])
	port := string(result[4])
	db_name := string(result[5])

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, username, password, db_name, port)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Printf("Database connect error:\n%s\n", err)
		return
	}

	fmt.Println("Database connected")
	db.AutoMigrate(&model.Pemeriksaan{}, &model.Penyakit{})
	fmt.Println("Migration Created")

	DB = db
}