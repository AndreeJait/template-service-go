package main

import (
	"context"
	"fmt"
	"github.com/AndreeJait/template-service-go/common/constant"
	"github.com/AndreeJait/template-service-go/config"
	"github.com/AndreeJait/go-utility/configw"
	"github.com/AndreeJait/go-utility/gracefull"
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/sqlw/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
)

type Api struct {
	log      loggerw.Logger
	cfg      *config.Config
	sqlW     postgres.SqlW
	pool     *pgxpool.Pool
	graceful *gracefull.GracefulShutDown
	e        *echo.Echo
}

func initServer(appMode constant.AppMode) {

	// init log
	var logFormat = loggerw.JSONFormatter
	if appMode == constant.Development {
		logFormat = loggerw.TextFormatter
	}
	log, err := loggerw.New(&loggerw.Option{
		Formatter: logFormat,
	})
	if err != nil {
		panic(err)
	}

	gracefully := gracefull.NewGracefulShutdown(log)

	// init config
	configInit := configw.New[config.Config](config.MapFiles,
		config.ConfigMode[string(appMode)])

	cfg, err := configInit.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Init SqlW - Repository
	pool, err := postgres.ConnectToDB(postgres.DbConfig{
		DBName:   cfg.DB.DBName,
		User:     cfg.DB.User,
		Host:     cfg.DB.Host,
		Password: cfg.DB.Password,
		Port:     cfg.DB.Port,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("connected to db")
	gracefully.AddFunc("shutdown-db-pool", func() error {
		pool.Close()
		return nil
	})
	sqlW := postgres.New(pool)

	// init echo
	e := echo.New()
	gracefully.AddFunc("shutdown-server", func() error {
		return e.Shutdown(context.Background())
	})
	var apiInstance = Api{
		log:      log,
		cfg:      cfg,
		sqlW:     sqlW,
		pool:     pool,
		e:        e,
		graceful: gracefully,
	}

	apiInstance.Start(appMode)
}

func (api *Api) Start(appMode constant.AppMode) {

	api.log.Infof("starting the server")
	s := make(chan os.Signal, 1)

	api.e.HideBanner = true
	api.e.HidePort = true

	api.e.Debug = appMode == constant.Development

	// init repository

	// init handler

	// init use case

	go func() {
		api.log.Infof("server start at %s", api.cfg.Service.Url)
		api.log.Error(api.e.Start(fmt.Sprintf("%s:%s",
			api.cfg.Service.Host, api.cfg.Service.Port)))
	}()

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	api.log.Infof("shutting down %s", api.cfg.Service.Name)
	// shutdown opened connection
	api.graceful.ShutdownAll()

	api.log.Infof("was shutdown gracefully")
}
