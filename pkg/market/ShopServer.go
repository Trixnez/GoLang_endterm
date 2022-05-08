package market

import (
	"context"
	"market/db"
	"market/pkg/api"
)

type ShopServer struct{}

func (*ShopServer) CreateShop(ctx context.Context, req *api.CreateShopRequest) (*api.ShopResponse, error) {
	db.CreateShop(db.CreateShopInput{
		Name: req.Name,
	})
	return &api.ShopResponse{}, nil
}

func (s *ShopServer) GetAllShops(ctx context.Context, req *api.EmptyInput) (*api.ShopsResponse, error) {
	shops := db.GetAllShops()
	var result []*api.ShopResponse
	for i := range shops {
		shop := shops[i]
		resultShop := &api.ShopResponse{
			Id:   uint32(shop.ID),
			Name: shop.Name,
		}

		result = append(result, resultShop)
	}
	return &api.ShopsResponse{Shops: result}, nil
}

func (s *ShopServer) GetShop(ctx context.Context, req *api.GetOneRequest) (*api.ShopResponse, error) {
	shop := db.GetShop(uint(req.GetId()))
	return &api.ShopResponse{Id: uint32(shop.ID), Name: shop.Name}, nil
}
