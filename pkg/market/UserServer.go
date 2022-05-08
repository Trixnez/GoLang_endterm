package market

import (
	"context"
	"market/db"
	"market/pkg/api"
)

type UserServer struct{}

func (*UserServer) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error) {
	db.CreateUser(db.CreateUserInput{
		Name:     req.Name,
		Surname:  req.Surname,
		Username: req.Username,
	})
	return &api.UserResponse{}, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, req *api.EmptyInput) (*api.UsersResponse, error) {
	users := db.GetAllUsers()
	var result []*api.UserResponse
	for i := range users {
		user := users[i]
		resultUser := &api.UserResponse{
			Id:       uint32(user.ID),
			Name:     user.Name,
			Surname:  user.Surname,
			Username: user.Username,
		}

		result = append(result, resultUser)
	}
	return &api.UsersResponse{Users: result}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *api.GetOneRequest) (*api.UserResponse, error) {
	user := db.GetUser(uint(req.GetId()))
	return &api.UserResponse{Id: uint32(user.ID), Name: user.Name, Surname: user.Surname, Username: user.Username}, nil
}
