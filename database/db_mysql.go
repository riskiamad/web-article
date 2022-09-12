package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/riskiamad/web-article/config"
)

var (
	DB  = dbConn()
	env = config.Config
)

// dbConn: Open connection for MySQL
func dbConn() *sql.DB {
	db, err := sql.Open("mysql", env.DbUser+":"+env.DbPass+"@tcp("+env.DbHost+")/"+env.DbName)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
