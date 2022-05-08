package market

import (
	"context"
	"market/db"
	"market/pkg/api"
)

type ProductServer struct{}

func (*ProductServer) CreateProduct(ctx context.Context, req *api.CreateProductRequest) (*api.ProductResponse, error) {
	db.CreateProduct(db.CreateProductInput{
		Name:   req.Name,
		Price:  req.Price,
		Amount: int(req.Amount),
		TypeID: uint(req.TypeID),
		ShopID: uint(req.ShopID),
	})
	return &api.ProductResponse{}, nil
}

func (s *ProductServer) GetAllProducts(ctx context.Context, req *api.EmptyInput) (*api.ProductsResponse, error) {
	products := db.GetAllProducts()
	var result []*api.ProductResponse
	for i := range products {
		product := products[i]
		resultProduct := &api.ProductResponse{
			Id:     uint32(product.ID),
			Name:   product.Name,
			Price:  product.Price,
			Amount: int32(product.Amount),
			TypeID: uint32(product.TypeID),
			ShopID: uint32(product.ShopID),
		}

		result = append(result, resultProduct)
	}
	return &api.ProductsResponse{Products: result}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, req *api.GetOneRequest) (*api.ProductResponse, error) {
	product := db.GetProduct(uint(req.GetId()))
	return &api.ProductResponse{Id: uint32(product.ID), Name: product.Name, Price: product.Price, Amount: int32(product.Amount), TypeID: uint32(product.TypeID), ShopID: uint32(product.ShopID)}, nil
}
