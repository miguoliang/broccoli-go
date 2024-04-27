package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserInfo struct {
	UserId   string `json:"sub"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUserInfoByContext(c *gin.Context, userInfo *UserInfo) error {

	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		return fmt.Errorf("no jwt token")
	}

	if gin.IsDebugging() {
		fmt.Println("jwtToken: ", jwtToken)
	}

	err := GetUserInfoByJwtToken(jwtToken, userInfo)
	if err != nil {
		return fmt.Errorf("no user id in token")
	}

	return nil
}

func GetUserInfoByJwtToken(jwtToken string, userInfo *UserInfo) error {

	token := jwtToken[7:]
	pieces := strings.Split(token, ".")
	if len(pieces) != 3 {
		return fmt.Errorf("invalid token")
	}

	payload := pieces[1]
	jsonStr, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		return err
	}

	var userData map[string]interface{}
	err = json.Unmarshal(jsonStr, &userData)
	if err != nil {
		return err
	}

	userInfo.UserId = userData["sub"].(string)
	userInfo.Username = userData["cognito:username"].(string)
	userInfo.Email = userData["email"].(string)

	return nil
}
