package responses

import (
	"go-chi-api/entity"
	web "go-chi-api/model"
)

func ConvertCreateUserResponse(user entity.User) web.CreateUserResponse {
	return web.CreateUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
}

func ConvertLoginUserResponse(token string) web.LoginUserResponse {
	return web.LoginUserResponse{
		Token: token,
	}
}

func ConvertUpdateUserResponse(user entity.User) web.UpdateUserResponse {
	return web.UpdateUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: *user.UpdatedAt,
	}
}
