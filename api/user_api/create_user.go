package user_api

import (
	"net/http"
	"time"

	"../../db"
	"../../utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

/**
创建用户
*/
func CreateUserHandler(c *gin.Context) {
	var user_request CreateUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	count, err := db.FindCount(db.DB, db.CollectionUser, bson.M{"username": user_request.Username})
	if err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "db err", Payload: nil})
		return
	}

	if count != 0 {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "user already exist", Payload: nil})
		return
	}

	user := User{
		UID:       utils.GenId(),
		Username:  user_request.Username,
		Password:  user_request.Password,
		OkexKey:   user_request.OkexKey,
		HuobiKey:  user_request.HuobiKey,
		Status:    USER_STATUS_ACTIVE,
		CreatedTS: time.Now(),
	}

	if err := db.Insert(db.DB, db.CollectionUser, user); err == nil {
		c.JSON(http.StatusOK, JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: nil})
	} else {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "db err", Payload: nil})
	}

}
