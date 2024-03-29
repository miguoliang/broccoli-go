package mock

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Client struct {
}

func (m *Client) VerifyToken(tokenStr string) (*jwt.Token, error) {
	return nil, nil
}

func (m *Client) Authorize(c *gin.Context) {
	c.Next()
}
