package database

// import (
// 	"fmt"
// 	"github.com/spf13/viper"
// 	"go_starter/logs"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// 	"log"
// 	"time"
// )

//type SqlLogger struct {
//	logger.Interface
//}
//

//var openMySQLConnectionDB *gorm.DB
//var errMySQL error
//
//func MysqlConnection() (*gorm.DB, error) {
//	myDSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
//		config.Env("mysql.host"),
//		config.Env("mysql.user"),
//		config.Env("mysql.password"),
//		config.Env("mysql.database"),
//		config.Env("mysql.port"),
//	)
//
//	fmt.Println("CONNECTING_TO_MYSQL_DB")
//	openMySQLConnectionDB, errMySQL = gorm.Open(mysql.Open(myDSN), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//		NowFunc: func() time.Time {
//			ti, _ := time.LoadLocation("Asia/Bangkok")
//			return time.Now().In(ti)
//		},
//	})
//	//DryRun: false,
//	if errMySQL != nil {
//		logs.Error(errMySQL)
//		log.Fatal("ERROR_PING_MYSQL", errMySQL)
//		return nil, errMySQL
//	}
//	fmt.Println("MYSQL_CONNECTED")
//	return openMySQLConnectionDB, nil
//}

// func MysqlConnection() (*gorm.DB, error) {
// 	myDSN := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
// 		viper.GetString("mysql.username"),
// 		viper.GetString("mysql.password"),
// 		viper.GetString("mysql.host"),
// 		viper.GetString("mysql.database"),
// 	)
// 	fmt.Println("CONNECTING_TO_DB_MYSQL")
// 	openConnectionDB, err := gorm.Open(mysql.Open(myDSN), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 		//DryRun: true,
// 		NowFunc: func() time.Time {
// 			tim, _ := time.LoadLocation("Asia/Bangkok")
// 			return time.Now().In(tim)
// 		},
// 	})
// 	if err != nil {
// 		logs.Error(err)
// 		log.Fatal("CONNECT_MYSQL_DATABASE_ERROR", err)
// 		return nil, err
// 	}

// 	fmt.Println("MYSQL_DB_CONNECTED")
// 	return openConnectionDB, nil
// }
