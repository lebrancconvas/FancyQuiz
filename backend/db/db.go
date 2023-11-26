package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gorp/gorp"
)

type DB struct {
	*sql.DB
}

var db *gorp.DbMap

func Init() {
	host := os.Getenv("DBHOST")

	port, err := strconv.Atoi(os.Getenv("DBPORT"))
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.PostgresDialect{},
	}

	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}
