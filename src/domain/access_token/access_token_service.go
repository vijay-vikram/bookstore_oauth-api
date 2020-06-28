package access_token

import (
	"github.com/vijay-vikram/bookstore_oauth-api/src/repository/services"
	"github.com/vijay-vikram/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	dbRepo        Repository
	restUsersRepo services.RestUsersRepository
}

func NewService(dbRepo Repository, restUsersRepo services.RestUsersRepository) Service {
	return &service{
		dbRepo:        dbRepo,
		restUsersRepo: restUsersRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token Id.")
	}

	return s.dbRepo.GetById(accessTokenId)
}

func (s *service) Create(atRequest AccessTokenRequest) (*AccessToken, *errors.RestErr) {
	if err := atRequest.Validate(); err != nil {
		return nil, err
	}

	//TODO:// Support both client_credentials and password grant type

	// Authenticate the user against user API:
	user, err := s.restUsersRepo.LoginUser(atRequest.Username, atRequest.Password)
	if err != nil {
		return nil, err
	}

	//Generate a new access token:
	token := GetNewAccessToken(user.Id)
	token.generate()

	//save the access token in cassandra db:
	if err = s.dbRepo.Create(token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
