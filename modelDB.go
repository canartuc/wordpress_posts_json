package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

// MySQLCon is the struct of values that needs to create MySQL connection
type MySQLCon struct {
	usr    string
	pass   string
	host   string
	port   string
	dbname string
}

// DB is global database object that used to be as connection pooling
var DB *sql.DB

// readConfig reads configuration values from environment variables and sets
func readConfig() *MySQLCon {
	config := &MySQLCon{
		usr:    os.Getenv("MYSQLDB_USR"),
		pass:   os.Getenv("MYSQLDB_PASS"),
		host:   os.Getenv("MYSQLDB_HOST"),
		port:   os.Getenv("MYSQLDB_PORT"),
		dbname: os.Getenv("MYSQLDB_DBNAME"),
	}

	return config
}

// DBCon controls and create database connection if there is not
func DBCon() *sql.DB {
	// If there is no pool/connection available, then let's open a new one
	if DB == nil {
		dbCfg := readConfig()

		// dbconnection string format: "username:password@tcp(127.0.0.1:3306)/test"
		DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbCfg.usr, dbCfg.pass,
			dbCfg.host, dbCfg.port, dbCfg.dbname))

		if err != nil {
			log.Fatal("(ERR) Database connection configuration is not satisfied: ", err)
		}

		// You can set the connection details here if you are using only this application for the database
		/*DB.SetMaxOpenConns(25)
		DB.SetMaxIdleConns(25)
		DB.SetConnMaxLifetime(5 * time.Minute)*/

		// In Go, connection doesn't mean that it is connected successfully so you need to ping the database if the
		// connection is really available with the given config
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		err = DB.PingContext(ctx)
		if err != nil {
			log.Fatalf("(ERR) Database: %s with host: %s is not reachable: %s", dbCfg.dbname, dbCfg.host,
				err.Error())
		}
	}

	return DB
}
