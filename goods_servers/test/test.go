package main

import (
	"Shop/goods_servers/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}

func GetBrandsTest() {
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name, " ", brand.Logo)
	}
}

func CreateBrandsTest() {
	rsp, err := brandClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: "中兴",
		Logo: "中兴的LOGO",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)

}

func DeleteBrandTest() {
	_, err := brandClient.DeleteBrand(context.Background(), &proto.BrandRequest{
		Name: "品牌198",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("删除成功")

}
func main() {
	Init()
	//GetBrandsTest()
	//CreateBrandsTest()
	//DeleteBrandTest()
}
