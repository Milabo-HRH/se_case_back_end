package main

import (
	"github.com/spf13/viper"
	"os"
	"se_case_back_end/common"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	//r := gin.Default()
	//r = routers.CollectRoute(r)
	//panic(r.Run("0.0.0.0:8080"))
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
