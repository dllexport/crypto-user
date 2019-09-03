package user_api

import (
	"net/http"

	"crypto-user/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

/**
设置APIkey
*/
func SetKeyUserHandler(c *gin.Context) {
	var user_request SetKeyUserRequest
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

	user.OkexKey = user_request.OkexKey
	user.HuobiKey = user_request.HuobiKey

	if err := db.Update(db.DB, db.CollectionUser, bson.M{"tel": user_request.Tel}, &user); err == nil {
		c.JSON(http.StatusOK, JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: nil})
	} else {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "update job status err", Payload: nil})
	}

}
