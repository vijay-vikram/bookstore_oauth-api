package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vijay-vikram/bookstore_oauth-api/src/domain/access_token"
	"github.com/vijay-vikram/bookstore_oauth-api/src/utils/errors"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, restErr := handler.service.GetById(c.Param("access_token_id"))

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var atRequest access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&atRequest); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	token, restErr := handler.service.Create(atRequest)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, token)
}
