package mysql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ClientConfig struct {
	Host     string
	Port     uint64
	Username string
	Password string
	Dbname   string
	Charset  string
	Pool     *ClientPoolConfig
}

type ClientPoolConfig struct {
	MaxIdleConns uint64
	MaxOpenConns uint64
	MaxLifetime  uint64
}

func NewClient(conf *ClientConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)", conf.Username, conf.Password, conf.Host, conf.Port)
	if conf.Dbname != "" {
		dsn = fmt.Sprintf("%s/%s", dsn, conf.Dbname)
	}
	dsn = fmt.Sprintf("%s?charset=%s&parseTime=true", dsn, conf.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	if conf.Pool != nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(int(conf.Pool.MaxIdleConns))
		sqlDB.SetMaxOpenConns(int(conf.Pool.MaxOpenConns))

		hours := time.Hour * time.Duration(conf.Pool.MaxLifetime)
		sqlDB.SetConnMaxLifetime(hours)
	}

	return db
}

var client *gorm.DB
