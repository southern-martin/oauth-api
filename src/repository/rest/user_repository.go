package rest

import (
	"encoding/json"
	"errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/southern-martin/oauth-api/src/domain/user"
	"github.com/southern-martin/utils-go/rest_error"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*user.User, rest_error.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*user.User, rest_error.RestErr) {
	request := user.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/user/login", request)

	if response == nil || response.Response == nil {
		return nil, rest_error.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_error.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_error.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	var user user.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_error.NewInternalServerError("error when trying to unmarshal user login response", errors.New("json parsing error"))
	}
	return &user, nil
}
