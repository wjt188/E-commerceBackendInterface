package main

import (
	"Shop/user_servers/model"
	password "Shop/user_servers/model/encode"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
)

// create database shop_user_srv CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
func genMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
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
	//造数据
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("123456", options)
	password_t := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("kol%d", i),
			Mobile:   fmt.Sprintf("1323433948%d", i),
			Password: password_t,
		}
		db.Save(&user)
	}
	//_ = db.AutoMigrate(&model.User{})

	//fmt.Println(genMD5("123456"))
	//result:e10adc3949ba59abbe56e057f20f883e
	//fmt.Println(genMD5("123457"))
	//salt, encodedPwd := password.Encode("generic password", nil)
	//fmt.Println("salt:", salt)
	//fmt.Println("encodepwd:", encodedPwd)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true
	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("123456", options)
	//password_t := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println("salt:", salt)
	//fmt.Println("encodepwd:", encodedPwd)
	//fmt.Println("password", password_t)
	//passwordInfo := strings.Split(password_t, "$")
	//check := password.Verify("123456", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check)
}
