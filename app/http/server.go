package main

import (
	"context"
	"fmt"
	"github.com/AndreeJait/go-utility/configw"
	"github.com/AndreeJait/go-utility/gracefull"
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/response"
	"github.com/AndreeJait/go-utility/sqlw/postgres"
	"github.com/AndreeJait/go-utility/timew"
	"github.com/AndreeJait/template-service-go/common/constant"
	"github.com/AndreeJait/template-service-go/config"
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
	appMode  constant.AppMode
	timeLoc  timew.Time
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

	timeLoc := timew.New(timew.LoadLocation("Asia/Jakarta"))
	var apiInstance = Api{
		log:      log,
		cfg:      cfg,
		sqlW:     sqlW,
		pool:     pool,
		e:        e,
		graceful: gracefully,
		appMode:  appMode,
		timeLoc:  timeLoc,
	}

	apiInstance.Start(appMode)
}

func (a *Api) Start(appMode constant.AppMode) {

	a.log.Infof("starting the server")
	s := make(chan os.Signal, 1)

	a.e.HideBanner = true
	a.e.HidePort = true

	a.e.Debug = appMode == constant.Development

	a.e.Use(loggerw.LoggerWitRequestID(a.log, a.e.Debug))

	a.e.HTTPErrorHandler = response.CustomHttpErrorHandler(a.log,
		response.MapDefaultErrResponse, a.e.Debug)

	a.startV1()

	go func() {
		a.log.Infof("server start at %s", a.cfg.Service.Url)
		a.log.Error(a.e.Start(fmt.Sprintf("%s:%s",
			a.cfg.Service.Host, a.cfg.Service.Port)))
	}()

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	a.log.Infof("shutting down %s", a.cfg.Service.Name)
	// shutdown opened connection
	a.graceful.ShutdownAll()

	a.log.Infof("was shutdown gracefully")
}
