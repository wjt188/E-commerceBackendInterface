package main

import (
	"Shop/user_servers/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {

	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 15,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		checkrsp, err := userClient.CheckPassWord(context.Background(), &proto.PassWordCheckInfo{
			PassWord:          "123456",
			EncryptedPassWord: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkrsp.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 5; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("kol%d", i),
			Mobile:   fmt.Sprintf("1016783624%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func TestGetUserByMobile() {

	for {
		var number string
		fmt.Println("请输入用户手机号：")
		fmt.Scanln(&number)
		if number == "quit" {
			fmt.Println("测试脚本已退出")
			break
		}
		rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
			Mobile: number,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp)

	}

}

func main() {
	Init()
	TestGetUserList()
	//TestCreateUser()
	//TestGetUserByMobile()
	conn.Close()
}
