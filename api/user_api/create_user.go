package user_api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"crypto-user/db"
	"crypto-user/utils"

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

	// username唯一
	count, err := db.FindCount(db.DB, db.CollectionUser, bson.M{"username": user_request.Username})
	if err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "db err", Payload: nil})
		return
	}

	if count != 0 {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "user already exist", Payload: nil})
		return
	}  

	// redis := utils.RedisUtils{}
	// redis.Connect()
	// defer redis.Close()

	// code, redisErr := redis.Get(user_request.Tel)
	// if redisErr != nil {
	// 	c.JSON(http.StatusInternalServerError, JSONReply{ErrorCode: -1, ErrorDescription: "code not found", Payload: nil})
	// 	return
	// }

	// if code != user_request.Code {
	// 	c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "code err", Payload: nil})
	// 	return
	// }

	//redis.Del(user_request.Tel)

	// add salt && sha256 password
	salt := utils.GenRandomStr(time.Now().UnixNano(), 64)
	h := sha256.New()
	h.Write([]byte(user_request.Password + salt))
	password := hex.EncodeToString(h.Sum(nil))

	fmt.Printf("%s", password)
	user := User{
		UID: utils.GenId(),
		//Tel:       user_request.Tel,
		Username:  user_request.Username,
		Password:  password,
		Salt:      salt,
		Status:    USER_STATUS_ACTIVE,
		CreatedTS: time.Now(),
	}

	if err := db.Insert(db.DB, db.CollectionUser, user); err == nil {
		c.JSON(http.StatusOK, JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: nil})
	} else {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "db err", Payload: nil})
	}

}
