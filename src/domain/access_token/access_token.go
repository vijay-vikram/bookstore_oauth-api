package access_token

import (
	"github.com/vijay-vikram/bookstore_oauth-api/src/utils/errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// Used for password gran_type
	Username string `json:"username"`
	Password string `json:"password"`
	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant type parameter")
	}

	//TODO://Validate parameters for each grant type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(time.Hour * expirationTime).Unix(),
	}
}

func (at *AccessToken) isExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token Id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("Invalid user Id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid client Id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid Expiration time")
	}
	return nil
}

func (at *AccessToken) generate() {
	var letterBytes = strconv.Itoa(int(at.UserId)) + strconv.Itoa(int(at.Expires))
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	at.AccessToken = string(b)
}
