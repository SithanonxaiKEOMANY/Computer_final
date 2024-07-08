package database

import (
	"fmt"
	
	"go_starter/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type SqlLogger struct {
	logger.Interface
}

var openPostgresConnectionDB *gorm.DB
var errPostgres error

func PostgresConnection() (*gorm.DB, error) {


	fmt.Println("CONNECTING_TO_POSTGRES_DB")
	openPostgresConnectionDB, errPostgres = gorm.Open(postgres.Open("postgresql://ceit_db:0urXtOt30QgI8d4ov5PFDbS07lvUSHqD@dpg-cq3pp9aju9rs739iep8g-a.singapore-postgres.render.com/ceit_db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Bangkok")
			return time.Now().In(ti)
		},
	})
	//DryRun: false,
	if errPostgres != nil {
		logs.Error(errPostgres)
		log.Fatal("ERROR_PING_POSTGRES", errPostgres)
		return nil, errPostgres
	}
	fmt.Println("POSTGRES_CONNECTED")
	return openPostgresConnectionDB, nil
}
