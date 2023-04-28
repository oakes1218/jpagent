package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// var ProductM *gorm.DB

type MysqlInit struct {
	ProductM *gorm.DB
}

func MysqlConn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=true", viper.Get("MYSQL_USER"), viper.Get("MYSQL_PASSWORD"), viper.Get("MYSQL_HOST"), viper.Get("MYSQL_PORT"), viper.Get("MYSQL_DATABASE"))
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	conn.DB().SetMaxIdleConns(viper.GetInt("MYSQL_MAXIDLE"))
	conn.DB().SetMaxOpenConns(viper.GetInt("MYSQL_MAXCONN"))
	conn.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("MYSQL_CONNMAXLIFETTIME")) * time.Second)
	conn.SingularTable(viper.GetBool("MYSQL_SINGULARTABLE"))
	conn.LogMode(viper.GetBool("MYSQL_LOGMODE"))
	conn.AutoMigrate(&Product{})

	return conn
}

func NewMysql() {
	Conn = &MysqlInit{ProductM: MysqlConn()}
}
