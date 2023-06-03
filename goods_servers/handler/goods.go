package handler

import "Shop/goods_servers/proto"

//GoodsList(context.Context, *GoodsFilterRequest) (*GoodsListResponse, error)
//BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
//CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
//DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
//UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
//GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

//func (s *GoodsServer) GoodsList(ctx context.Context, client *proto.GoodsFilterRequest) (proto.GoodsListResponse, error) {
//
//}
