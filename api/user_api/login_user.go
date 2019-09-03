package user_api

import (
	"fmt"
	"net/http"
	"time"

	"crypto-user/utils"

	"crypto-user/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/**
启动作业
*/
func LoginUserHandler(c *gin.Context) {
	var user_request LoginUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	var user User
	if err := db.FindOne(db.DB, db.CollectionUser, bson.M{"tel": user_request.Tel}, nil, &user); err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	// if password error
	if !checkPassword(user_request.Password, user.Password, user.Salt) {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "user password incorrect", Payload: nil})
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
		c.JSON(http.StatusInternalServerError, JSONReply{ErrorCode: -1, ErrorDescription: "jwt sign err", Payload: nil})
		return
	}

	c.JSON(http.StatusOK, JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}})

}
