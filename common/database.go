package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"se_case_back_end/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("dataSource.driverName")
	host := viper.GetString("dataSource.host")
	port := viper.GetString("dataSource.port")
	database := viper.GetString("dataSource.database")
	username := viper.GetString("dataSource.username")
	password := viper.GetString("dataSource.password")
	charset := viper.GetString("dataSource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.Register{}, &model.Case{}, &model.Supplement{}, &model.Treatment{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
