package db

import (
	"fmt"
	"log"

	"BlogsAPI/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	DBCon *gorm.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func CreateDatabase() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("DB Connection failed")
	}

	migrateDatabase(db)
	return db, nil
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Blog{})
	db.AutoMigrate(&models.Admin{})
}
