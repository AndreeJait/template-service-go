package main

import (
	"database/sql"
	"fmt"
	"github.com/AndreeJait/exam-platform-service/common/constant"
	"github.com/AndreeJait/exam-platform-service/config"
	"github.com/AndreeJait/go-utility/configw"
	"github.com/AndreeJait/go-utility/loggerw"
	migrate "github.com/rubenv/sql-migrate"

	_ "github.com/lib/pq"
)

func StartMigrate(migrateType constant.MigrateType, appMode constant.AppMode) {

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

	cfgInit := configw.New[config.Config](config.MapFiles, config.Development)

	cfg, err := cfgInit.LoadConfig()
	if err != nil {
		log.Error(err)
		return
	}

	connection := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
	)
	sqldb, err := sql.Open("postgres", connection)
	if err != nil {
		log.Error(err)
		return
	}

	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("[%s] start process migrate %s", string(appMode), string(migrateType))

	if appMode == constant.Production && migrateType == constant.MigrateTypeFresh {
		log.Infof("can't run fresh in production")
		return
	} else {
		log.Infof("drop database")
		_, err = sqldb.Exec(`DROP SCHEMA public CASCADE; CREATE SCHEMA public;`)
		if err != nil {
			panic(err)
		}
		sqldb, err = sql.Open("postgres", connection)
		if err != nil {
			log.Error(err)
			return
		}
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./files/migrations/postgres",
	}

	var direction migrate.MigrationDirection
	switch migrateType {
	case constant.MigrateTypeUp:
		direction = migrate.Up
	case constant.MigrateTypeDown:
		direction = migrate.Down
	case constant.MigrateTypeFresh:
		direction = migrate.Up
	}

	count, err := migrate.Exec(sqldb, "postgres", migrations, direction)
	if err != nil {
		panic(err)
	}

	log.Infof("done to execute %d migrate file", count)
}
