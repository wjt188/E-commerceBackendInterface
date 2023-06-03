package initialize

import (
	"Shop/goods_servers/global"
	"fmt"
	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
func InitConfig() {
	debug := GetEnvInfo("SHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("goods_servers/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("goods_servers/%s-debug.yaml", configFilePrefix)
	}
	fmt.Println(configFileName)
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
