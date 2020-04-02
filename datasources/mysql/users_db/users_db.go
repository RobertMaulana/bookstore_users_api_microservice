package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//const (
//	mysql_users_username = "database_users_username"
//	mysql_users_password = "database_users_password"
//	mysql_users_host = "database_users_host"
//	mysql_users_schema = "database_users_schema"
//)

var (
	Client *sql.DB

	//username = os.Getenv(mysql_users_username)
	//password = os.Getenv(mysql_users_password)
	//host = os.Getenv(mysql_users_host)
	//schema = os.Getenv(mysql_users_schema)

	username = "root"
	password = "root"
	host = "127.0.0.1:3306"
	schema = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil{
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}