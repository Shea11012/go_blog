package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shea11012/go_blog/cmd"
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/internal/models"
	"github.com/shea11012/go_blog/internal/routers"
	"github.com/shea11012/go_blog/pkg/logger"
	"github.com/shea11012/go_blog/pkg/setting"
	"github.com/shea11012/go_blog/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port string
	runMode string
	config string
	isVersion bool
	buildTime string
	buildVersion string
	gitCommitID string
)

func init() {
	setupFlag()
	loadSetting()
	setupTracer()
	setupDBEngine()
	setupLogger()
}

// @title 博客 API
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	if isVersion {
		fmt.Printf("build_time: %s\n",buildTime)
		fmt.Printf("build_version: %s\n",buildVersion)
		fmt.Printf("git_commit_id: %s\n",gitCommitID)
		return
	}
	gin.SetMode(global.ServerSetting.RunMode)
	err := cmd.Execute()
	global.CheckError(err)

	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err = s.ListenAndServe()
		global.CheckError(err)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:",err)
	}

	log.Println("Server exiting")
}

func loadSetting() {
	settings,err := setting.NewSetting(strings.Split(config,",")...)
	global.CheckError(err)

	err = settings.ReadSection("Server",&global.ServerSetting)
	global.CheckError(err)

	err = settings.ReadSection("App",&global.AppSetting)
	global.CheckError(err)

	err = settings.ReadSection("Database",&global.DatabaseSetting)
	global.CheckError(err)

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = settings.ReadSection("JWT", &global.JWTSetting)
	global.CheckError(err)

	global.JWTSetting.Expire *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}

	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
}

func setupDBEngine() {
	var err error
	global.DBEngine,err = models.NewDBEngine(global.DatabaseSetting)
	global.CheckError(err)
}

func setupLogger() {
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: filename,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
}

func setupTracer() {
	jaegerTracer,_,err := tracer.NewJaegerTracer("blog-service","127.0.0.1:6831")
	if err != nil {
		global.CheckError(err)
	}

	global.Tracer = jaegerTracer
}

func setupFlag()  {
	flag.StringVar(&port,"port","","启动端口")
	flag.StringVar(&runMode,"mode","","启动模式")
	flag.StringVar(&config,"config","configs/","指定配置文件路径")
	flag.BoolVar(&isVersion,"version",false,"编译信息")
	flag.Parse()
}
