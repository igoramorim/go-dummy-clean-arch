package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DBConnections struct {
	MySQL *sql.DB
}

func OpenConnections() (*DBConnections, error) {
	mySQLConnection, err := openMySQLConnection()
	if err != nil {
		return nil, err
	}

	return &DBConnections{MySQL: mySQLConnection}, nil
}

func openMySQLConnection() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)

	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		return nil, err
	}

	log.Printf("Connected to host: %s user: %s db: %s", dbHost, dbUsername, dbName)

	return DB, nil
}
