package pkg

import (
	"database/sql"
	"log"
	"os"

	"github.com/RipulHandoo/blogx/authentication/db/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

)

func DbInstance() *database.Queries {
	// load the .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	// get the database connectino string
	db_url := os.Getenv("DB_URL")

	if db_url == "" {
		log.Fatal("DB_URL not set in .env file")
	}

	// setup the databse connection
	db, dbErr := sql.Open("postgres", db_url)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	DbQueries := database.New(db)
	return DbQueries
}

var DbClient *database.Queries = DbInstance()