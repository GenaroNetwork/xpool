package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

const defaultMaxConnections = 20

var (
	connection *gorm.DB
)

func init() {
	connection = connect()
}

func GetDB() *gorm.DB {
	return connection
}

const DATABASEURL  = "root:@/xpool?charset=utf8&parseTime=True&loc=Local"
func connect() *gorm.DB {
	max := getMaxConnection()
	conn, err := gorm.Open("mysql", DATABASEURL)
	if err != nil {
		panic(err)
	}

	conn.DB().SetMaxIdleConns(max / 5)
	conn.DB().SetMaxOpenConns(max)
	conn.SingularTable(true)
	conn.LogMode(true)
	return conn
}

func getMaxConnection() int {
	env := os.Getenv("DATABASE_MAX_CONNECTIONS")
	if env == "" {
		return defaultMaxConnections
	}

	max, _ := strconv.Atoi(env)
	return max
}
