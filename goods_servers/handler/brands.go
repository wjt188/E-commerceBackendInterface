package handler

import (
	"Shop/goods_servers/global"
	"Shop/goods_servers/model"
	"Shop/goods_servers/proto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//BrandList(context.Context, *BrandFilterRequest) (*BrandListResponse, error)
//CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error)
//DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
//UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)

func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}
	var brands []model.Brands
	//result := global.DB.Find(&brands)
	//if result.Error != nil {
	//	return nil, result.Error
	//}

	//处理品牌数据分页逻辑
	result := global.DB.Scopes(Paginate(2, 10)).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}
	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)
	var brandResponses []*proto.BrandInfoResponse
	for _, brand := range brands {
		//第1种写法
		//brandResponse := proto.BrandInfoResponse{
		//	Id:   brand.ID,
		//	Name: brand.Name,
		//	Logo: brand.Logo,
		//}
		//brandsResponses = append(brandsResponses, &brandResponse)
		//第二种写法
		brandResponses = append(brandResponses, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}

func (s *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	//先判断品牌是否存在，若存在，则返回错误
	//if result := global.DB.First(&model.Brands{Name: req.Name}); result.RowsAffected == 1 {
	//	return nil, status.Errorf(codes.InvalidArgument, "该品牌已存在")
	//}
	var brand model.Brands
	result := global.DB.Where(&model.Brands{Name: req.Name}).First(&brand)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "该品牌已存在")
	}

	//示例华一个对象，从请求中拿到相应参数，保存到该对象中
	brands := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	//保存在数据库中
	global.DB.Save(brands)
	return &proto.BrandInfoResponse{Id: brands.ID}, nil
}

func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	//if result := global.DB.Delete(&model.Brands{Name: req.Name}); result.RowsAffected == 0 {
	//	return nil, status.Errorf(codes.NotFound, "该品牌不存在")
	//}
	var brands model.Brands
	result := global.DB.Where(&model.Brands{Name: req.Name}).First(&brands)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "该品牌不存在")
	}
	fmt.Println(brands.ID)

	global.DB.Unscoped().Delete(&brands)
	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brands := model.Brands{}
	if result := global.DB.First(&model.Brands{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "该品牌不存在")
	}
	if req.Name != "" {
		brands.Name = req.Name
	}
	if req.Logo != "" {
		brands.Logo = req.Logo
	}
	global.DB.Save(&brands)
	return &emptypb.Empty{}, nil
}
