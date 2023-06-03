package main

import (
	"Shop/goods_servers/proto"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var brandClient2 proto.GoodsClient
var conn2 *grpc.ClientConn

func Init2() {
	var err error
	conn2, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient2 = proto.NewGoodsClient(conn2)
}
func GetBannerListTest() {
	rsp, err := brandClient2.BannerList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	for _, banner := range rsp.Data {
		fmt.Println(banner.Id, " ", banner.Image, " ", banner.Url, " ", banner.Index)
	}
}

func main() {
	Init2()
	GetBannerListTest()
}
