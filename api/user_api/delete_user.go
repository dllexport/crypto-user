package user_api

import (
	"net/http"

	"../../db"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

/**
启动作业
*/
func DeleteUserHandler(c *gin.Context) {
	var user_request DeleteUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	var user User
	if err := db.FindOneById(db.DB, db.CollectionUser, user_request.UID, &user); err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	user.Status = USER_STATUS_DELETED
	if err := db.Update(db.DB, db.CollectionUser, bson.M{"_id": user_request.UID}, &user); err == nil {
		c.JSON(http.StatusOK, JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: nil})
	} else {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "update job status err", Payload: nil})
	}

}
