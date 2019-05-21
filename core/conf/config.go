package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config") // 设置配置文件名 (不带后缀)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")    // 第一个搜索路径
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
