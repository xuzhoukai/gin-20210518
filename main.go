package main

import (
	mgo "gin-20210518/dao/mongo"
	"gin-20210518/routers"
	"gin-20210518/routers/geet"
	"gin-20210518/routers/student"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"os"
)

var log = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.Formatter = &logrus.JSONFormatter{}
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//f, _ := os.Create("./gin.log")
	f, _ := os.OpenFile("./gin.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend|os.ModePerm)
	log.Out = f
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	// Only log the warning severity or above.
	log.Level = logrus.InfoLevel
	log.SetReportCaller(true)
}
func main() {
	//// 创建一个默认的路由引擎
	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//	time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})

	//router.GET("/hello")

	routers.Include(geet.Routers, student.Routers)
	router := routers.Init()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 优雅的停服务器
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelFunc()
		if err := srv.Shutdown(ctx); err != nil {
			logrus.Fatalf("fial to close server err:%v", err)
		}
		log.Println("server gracefully shutdown")
		close(processed)
	}()
	// 开启一个goroutine启动服务
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	<-processed
}

func initMongodb() {
	uri := ""
	dbName := ""
	timeOut := 5 * time.Second
	var maxConnect uint64 = 100

	_, err := mgo.ConnectToDB(uri, dbName, timeOut, maxConnect)
	if err != nil {
		log.Fatalf("connect to mongodb err:%v", err)
	}
}
