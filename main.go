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
	"github.com/spf13/viper"
)

var Validate *validator.Validate = validator.New()
var Val *Config

type Config struct {
	ServerPort          string `mapstructure:"SERVER_PORT" json:"SERVER_PORT"`
	MysqlHost           string `mapstructure:"MYSQL_HOST" json:"MYSQL_HOST"`
	MysqlPort           int    `mapstructure:"MYSQL_PORT" json:"MYSQL_PORT"`
	MysqlUser           string `mapstructure:"MYSQL_USER" json:"MYSQL_USER"`
	MysqlPassword       string `mapstructure:"MYSQL_PASSWORD" json:"MYSQL_PASSWORD"`
	MysqlMaxidle        int    `mapstructure:"MYSQL_MAXIDLE" json:"MYSQL_MAXIDLE"`
	MysqlMaxconn        int    `mapstructure:"MYSQL_MAXCONN" json:"MYSQL_MAXCONN"`
	MysqlConnMaxLifeTim int    `mapstructure:"MYSQL_CONNMAXLIFETTIME" json:"MYSQL_CONNMAXLIFETTIME"`
	MysqlSingularTable  bool   `mapstructure:"MYSQL_SINGULARTABLE" json:"MYSQL_SINGULARTABLE"`
	MysqlLogMode        bool   `mapstructure:"MYSQL_LOGMODE" json:"MYSQL_LOGMODE"`
}

func init() {
	//讀取還近變數
	viper.AutomaticEnv()
	viper.SetEnvPrefix("JP")

	//取設定黨
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&Val)
	if err != nil {
		panic(err)
	}

	log.Println("ENV:", Val)
	log.Println("Cofing 設定成功")
	//建立mysql pool
	model.NewMysql()
	//初始log框架
	jplog.LogInit()
}

func main() {
	WaitShutdown(func() {
		if err := model.Conn.Close(); err != nil {
			log.Println(err)
			return
		}

		log.Println("3秒後服務停止")
		time.Sleep(3 * time.Second)
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
	r.Run(string(fmt.Sprintf("%v", viper.Get("SERVER_PORT"))))
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func MiddleValid(c *gin.Context) {
	c.Next()
}
