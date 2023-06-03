package main

import (
	"Shop/goods_servers/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/shop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newlogger,
	})
	if err != nil {
		panic(err)
	}

	for i := 1; i < 10; i++ {
		banners := model.Banner{
			Image: fmt.Sprintf("image%d", i),
			Url:   fmt.Sprintf("url%d", i),
			Index: int32(i),
		}
		db.Save(&banners)
	}

}
