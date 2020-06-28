package services

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/vijay-vikram/bookstore_oauth-api/src/domain/users"
	"github.com/vijay-vikram/bookstore_oauth-api/src/utils/errors"
	"time"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	loginRequest := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := restClient.Post("/users/login", loginRequest)
	// Timeout situation
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest client response when try to login user.")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to login user.")
		}
		return nil, &restErr
	}

	var user users.User
	err := json.Unmarshal(response.Bytes(), &user)
	if err != nil {
		return nil, errors.NewInternalServerError("Error when trying to unmarshal users response")
	}
	return &user, nil
}
