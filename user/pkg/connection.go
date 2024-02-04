package pkg

import (
	"database/sql"
	"log"
	"os"

	"github.com/RipulHandoo/blogx/user/db/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DbInstane() *database.Queries {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file in user/pkg/connection.go ERROR: %v", err)
	}

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("DB_URL not set in .env file")
	}

	db, dbErr := sql.Open("postgres", db_url)
	if dbErr != nil {
		log.Fatal("Error in connection.go: %v" ,dbErr)
	}

	DbQueries := database.New(db)
	return DbQueries
}

var DbClient *database.Queries = DbInstane()
