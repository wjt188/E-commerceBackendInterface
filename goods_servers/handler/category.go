package handler

import (
	"Shop/goods_servers/global"
	"Shop/goods_servers/model"
	"Shop/goods_servers/proto"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

//GetAllCategorysList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoryListResponse, error)
//// 获取子分类
//GetSubCategory(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*SubCategoryListResponse, error)
//CreateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*CategoryInfoResponse, error)
//DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
//UpdateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)

//实现获取数据库中三级目录的数据
func (s *GoodsServer) GetAllCategorysList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*proto.CategoryListResponse, error) {
	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory ").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}
