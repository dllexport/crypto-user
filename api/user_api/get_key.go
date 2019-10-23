package user_api

import (
	"net/http"

	"crypto-user/api"
	"crypto-user/db"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type APIKEY struct {
	OKEX_KEY  OkexKeyDetail  `bson:"okex_key" json:"okex_key"`
	HUOBI_KEY HuobiKeyDetail `bson:"huobi_key" json:"huobi_key"`
	PUSH_URL  string         `bson:"push_url" json:"push_url"`
}

/**
设置APIkey
*/
func GetKeyUserHandler(c *gin.Context) {

	claims := jwt.ExtractClaims(c)

	var user User
	if err := db.FindOneById(db.DB, db.CollectionUser, claims["uid"], &user); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	key := APIKEY{
		OKEX_KEY:  user.OkexKey,
		HUOBI_KEY: user.HuobiKey,
		PUSH_URL:  user.PushURL,
	}

	c.JSON(http.StatusOK, api.JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: key})
}
