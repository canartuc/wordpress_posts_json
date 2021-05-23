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

type Post struct {
	PostID     int    `json:"post_id"`
	PostTitle  string `json:"post_title"`
	PostUrl    string `json:"post_url"`
	PostPoster string `json:"post_poster"`
}

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
	dbCfg := readConfig()

	// dbconnection string format: "username:password@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbCfg.usr, dbCfg.pass,
		dbCfg.host, dbCfg.port, dbCfg.dbname))

	if err != nil {
		log.Fatal("(ERR) Database connection configuration is not satisfied: ", err)
	}

	// You can set the connection details here if you are using only this application for the database
	/*db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)*/

	// In Go, connection doesn't mean that it is connected successfully so you need to ping the database if the
	// connection is really available with the given config
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("(ERR) Database: %s with host: %s is not reachable: %s", dbCfg.dbname, dbCfg.host,
			err.Error())
	}

	return db

}

func GetAllPosts() ([]Post, error) {
	var res []Post
	db := DBCon()

	// Execute the query which takes all posts from Wordpress database and order by New to Old
	results, err := db.Query("SELECT post.ID as post_id, " +
		"post.post_title as post_title, post.guid as post_url, p.guid as post_poster " +
		"FROM wp_posts AS post " +
		"JOIN wp_postmeta AS postmeta ON postmeta.post_id = post.ID " +
		"JOIN wp_posts AS p ON p.ID = postmeta.meta_value " +
		"WHERE post.post_type = 'post' AND post.post_status = 'publish' AND postmeta.meta_key = '_thumbnail_id' " +
		"ORDER BY post.post_modified DESC;")
	if err != nil {
		log.Fatal("(ERR) Cannot query the database: ", err)
	}

	// Iterate through results
	for results.Next() {
		var eachRes = new(Post)
		// for each row of the data, scan into structure
		err = results.Scan(&eachRes.PostID, &eachRes.PostTitle, &eachRes.PostUrl, &eachRes.PostPoster)
		if err != nil {
			log.Println("(ERR) Problem while reading data from database: ", err)
		}

		res = append(res, *eachRes)
	}
	defer db.Close()

	return res, err
}
