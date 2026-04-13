package service

import (
	"context"
	"project/internal/model"
	"project/internal/repository"
	"sync"
)

type UserResponse struct {
	User    model.User `json:"user"`
	Profile string     `json:"profile"`
	Orders  string     `json:"orders"`
}

func GetUserFullData(ctx context.Context, userID int64) (UserResponse, error) {

	var wg sync.WaitGroup
	var user model.User
	var profile string
	var orders string
	var err error

	wg.Add(3)

	go func() {
		defer wg.Done()

		u, e := repository.GetUserByID(ctx, userID)

		if e != nil {
			err = e
			return
		}

		user = u
	}()

	go func() {
		defer wg.Done()
		profile = "profile-data"
	}()

	go func() {
		defer wg.Done()
		orders = "orders-data"
	}()

	wg.Wait()

	return UserResponse{
		User:    user,
		Profile: profile,
		Orders:  orders,
	}, err
}
