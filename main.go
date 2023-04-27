package main

import (
	"context"
	"fmt"
	jplog "jpagent/log"
	"jpagent/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var Validate *validator.Validate = validator.New()

func init() {
	//取設定黨
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//建立mysql pool
	model.ProductM = mysqlConn()
	//初始log框架
	jplog.LogInit()
}

func main() {
	WaitShutdown(func() {
		if err := model.ProductM.Close(); err != nil {
			log.Println(err)
			return
		}

		// time.Sleep(3 * time.Second)
	})
	httpServer()
}

func WaitShutdown(callback func()) context.Context {
	//grace shutdown
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)

	go func() {
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		defer signal.Stop(c)

		<-c
		fmt.Println("[notice] get cancel signal")
		cancel()
		callback()
		os.Exit(0)
	}()

	return ctx
}

func PanicRecover(err *error) {
	if r := recover(); r != nil {
		log.Println("PanicRecover")
	}
}

func httpServer() {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(MiddleValid)

	r.GET("/ping", Pong)

	prefix := r.Group("/quote")
	prefix.POST("/insert", CreateQuote)
	prefix.GET("/", GetQuote)
	prefix.DELETE("/:id", DeleteQuote)
	prefix.PUT("/update", UpdateQuote)
	r.Run(string(fmt.Sprintf("%v", viper.Get("JP_SERVER_PORT"))))
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func MiddleValid(c *gin.Context) {
	c.Next()
}

func mysqlConn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", viper.Get("JP_MYSQL_USER"), viper.Get("JP_MYSQL_PASSWORD"), viper.Get("JP_MYSQL_HOST"), viper.Get("JP_MYSQL_PORT"), viper.Get("JP_MYSQL_DATABASE"))
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	conn.DB().SetMaxIdleConns(viper.GetInt("JP_MYSQL_MAXIDLE"))
	conn.DB().SetMaxOpenConns(viper.GetInt("JP_MYSQL_MAXCONN"))
	conn.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("JP_MYSQL_CONNMAXLIFETTIME")) * time.Second)
	conn.SingularTable(viper.GetBool("JP_MYSQL_SINGULARTABLE"))
	conn.LogMode(viper.GetBool("JP_MYSQL_LOGMODE"))
	conn.AutoMigrate(&model.Product{})

	return conn
}
