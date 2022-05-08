package main

import (
	"google.golang.org/grpc"
	"log"
	"market/db"
	"market/pkg/api"
	"market/pkg/market"
	"net"
)

func main() {
	db.ConnectDB()
	s := grpc.NewServer()
	userServer := &market.UserServer{}
	api.RegisterUserServer(s, userServer)
	cardServer := &market.CardServer{}
	api.RegisterCardServer(s, cardServer)
	typeServer := &market.TypeServer{}
	api.RegisterTypeServer(s, typeServer)
	shopServer := &market.ShopServer{}
	api.RegisterShopServer(s, shopServer)
	productServer := &market.ProductServer{}
	api.RegisterProductServer(s, productServer)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
