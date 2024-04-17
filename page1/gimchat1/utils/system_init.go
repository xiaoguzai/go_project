package utils

import (
	"fmt"
	"gimchat1/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
}
func InitMySQL() (*gorm.DB, error) {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	//DB.Find(model.UnderBasic{})
	user := model.UserBasic{}
	DB.Find(&user)
	fmt.Println("user:", user)
	return nil, nil
}
