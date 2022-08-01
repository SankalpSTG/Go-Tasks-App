package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ServerConfig struct{
	Host string `json:"host"`
	Port int `json:"port"`
	User string `json:"mysql-user"`
	Password string `json:"mysql-password"`
	DatabaseName string `json:"db-name"`
}
func getServerConfig() ServerConfig{
	return ServerConfig{"localhost", 3306, "root", "", "tasks"}
}
func GetDatabase() (*sql.DB, error){
	config := getServerConfig()
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.User, config.Password, config.Host, config.Port, config.DatabaseName))
	if(err == nil){
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		err = db.Ping()
	}
	return db, err
}