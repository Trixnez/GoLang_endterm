package market

import (
	"context"
	"market/db"
	"market/pkg/api"
)

type TypeServer struct{}

func (*TypeServer) CreateType(ctx context.Context, req *api.CreateTypeRequest) (*api.TypeResponse, error) {
	db.CreateType(db.CreateTypeInput{
		Name: req.Name,
	})
	return &api.TypeResponse{}, nil
}

func (s *TypeServer) GetAllTypes(ctx context.Context, req *api.EmptyInput) (*api.TypesResponse, error) {
	types := db.GetAllTypes()
	var result []*api.TypeResponse
	for i := range types {
		_type := types[i]
		resultType := &api.TypeResponse{
			Id:   uint32(_type.ID),
			Name: _type.Name,
		}

		result = append(result, resultType)
	}
	return &api.TypesResponse{Types: result}, nil
}

func (s *TypeServer) GetType(ctx context.Context, req *api.GetOneRequest) (*api.TypeResponse, error) {
	_type := db.GetType(uint(req.GetId()))
	return &api.TypeResponse{Id: uint32(_type.ID), Name: _type.Name}, nil
}
