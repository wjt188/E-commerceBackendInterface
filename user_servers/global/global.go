package global

import (
	"Shop/user_servers/config"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)

//func CheckError(err error){
//	if err != nil{
//		panic(err)
//	}
//}
//
//func InitDB() *gorm.DB {
//	user := ServerConfig.MysqlInfo.User
//	password := ServerConfig.MysqlInfo.Password
//	host := ServerConfig.MysqlInfo.Host
//	port := ServerConfig.MysqlInfo.Port
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=local",user,password,host,port)
//
//
//	DB , err := gorm.Open(mysql.Open(dsn),&gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//	})
//	CheckError(err)
//	return DB
//}

//
//func Init() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
//			SlowThreshold: time.Second,
//			LogLevel:      logger.Info,
//			Colorful:      true,
//		},
//	)
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//		Logger: newLogger,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//}
