package market

import (
	"context"
	"market/db"
	"market/pkg/api"
)

type CardServer struct{}

func (*CardServer) CreateCard(ctx context.Context, req *api.CreateCardRequest) (*api.CardResponse, error) {
	db.CreateCard(db.CreateCardInput{
		Balance: req.Balance,
		UserId:  uint(req.UserId),
	})
	return &api.CardResponse{}, nil
}

func (s *CardServer) GetCard(ctx context.Context, req *api.GetOneRequest) (*api.CardResponse, error) {
	card := db.GetCard(uint(req.GetId()))
	return &api.CardResponse{Id: uint32(card.ID), Balance: card.Balance, UserId: uint32(card.UserID)}, nil
}
