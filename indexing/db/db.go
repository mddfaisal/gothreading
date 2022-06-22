package db

import (
	"database/sql"
	"indexing/utils"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
)

// global db connection variables
var (
	db   *sql.DB
	once sync.Once
)

// db connectivity
func Db() *sql.DB {
	once.Do(func() {
		dsn := mysql.Config{
			User:                 os.Getenv("DB_USER"),
			Passwd:               os.Getenv("DB_PASS"),
			Net:                  os.Getenv("DB_NET"),
			Addr:                 os.Getenv("DB_ADDR"),
			DBName:               os.Getenv("DB_NAME"),
			AllowNativePasswords: true,
		}
		conn, err := sql.Open(os.Getenv("DB_DRIVER"), dsn.FormatDSN())
		db = conn
		utils.Panic(err)
	})
	return db
}
