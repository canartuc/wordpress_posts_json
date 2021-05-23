# wordpress_posts_json

## General
This project creates JSON representation of WordPress posts directly from MySQL database.
To make it run, you need to have the following **environment variables:**
- MYSQLDB_USR : mysql connection username (read-only user is more than enough)
- MYSQLDB_PASS : mysql connection password
- MYSQLDB_HOST : host address of the mysql (for example: localhost, 127.0.0.1)
- MYSQLDB_PORT : port of the mysql database (it is not set to 3306 by default!)
- MYSQLDB_DBNAME : the database you would like to use (WordPress database)

If you are not happy with the environment variables' name, you can modify them in `modelDB.go` 
inside `readConfig()` function.

## Installation
You can use release page for the different platforms and environments but if you would like to build yourself,
this is a very generic go code, you don't need to take extra care. Here are the steps:
1. If you don't have Go installed, please follow: https://golang.org/doc/install
2. Clone the repo: `git clone git@github.com:canartuc/wordpress_posts_json.git`
3. Inside the folder: `go mod tidy`
4. Then, compile: `go build -o __output_name__ .`

## Running
Depends on the platform, you need to run executable/binary file with command line parameter `out`.
Here is the sample for direct binary of macOS and Linux:
```sh
./wordpress_posts_json -out=/home/username/Desktop/output.json
```
Then you can see the output of it inside the path you defined. If you would like to make the output
json file reachable via web, you could define the default folder of your web server (Apache, NGINX etc.)