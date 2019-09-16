package user_api

import (
	"fmt"
	"net/http"

	"crypto-user/api"
	"crypto-user/db"

	"gopkg.in/mgo.v2/bson"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/**
设置APIkey
*/
func SetKeyUserHandler(c *gin.Context) {
	var user_request SetKeyUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	claims := jwt.ExtractClaims(c)
	
	var user User
	if err := db.FindOneById(db.DB, db.CollectionUser, claims["uid"], &user); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	user.OkexKey = user_request.OkexKey
	user.HuobiKey = user_request.HuobiKey

	if err := db.Update(db.DB, db.CollectionUser, bson.M{"_id": claims["uid"]}, &user); err == nil {
		c.JSON(http.StatusOK, api.JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: nil})
	} else {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "update job status err", Payload: nil})
	}

}
