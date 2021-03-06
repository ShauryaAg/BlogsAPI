package db

import (
	"fmt"
	"log"
	"os"

	"BlogsAPI/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	DBCon *gorm.DB
)

var (
	host     = os.Getenv("POSTGRES_HOST")
	port     = 5432
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

func CreateDatabase() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("DB Connection failed: ", err)
	}

	migrateDatabase(db)
	return db, nil
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Blog{})
	db.AutoMigrate(&models.Admin{})
}
