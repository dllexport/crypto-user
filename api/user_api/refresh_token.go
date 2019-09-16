package user_api

import (
	"fmt"
	"net/http"
	"time"

	"crypto-user/api"
	"crypto-user/utils"

	"crypto-user/db"

	jwt_gin "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

/**
启动作业
*/
func RefreshTokenHandler(c *gin.Context) {

	claims := jwt_gin.ExtractClaims(c)

	var user User
	if err := db.FindOneById(db.DB, db.CollectionUser, claims["uid"], &user); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	secretKey, _ := utils.GetConfig().Get("jwt.secret")
	expireTime, _ := utils.GetConfig().GetInt("jwt.expire_time")

	// sign jwt and reply
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":            user.UID,
		"allow_strategy": user.AllowStrategy,
		"exp":            time.Now().Local().Add(time.Hour * time.Duration(expireTime)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(tokenString, err)
		c.JSON(http.StatusInternalServerError, api.JSONReply{ErrorCode: -1, ErrorDescription: "jwt sign err", Payload: nil})
		return
	}

	c.JSON(http.StatusOK, api.JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}})

}
