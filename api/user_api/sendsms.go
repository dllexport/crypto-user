package user_api

import (
	"fmt"
	"net/http"
	"time"

	"crypto-user/api"
	"crypto-user/db"
	"crypto-user/utils"

	"github.com/gin-gonic/gin"
	"github.com/qinxin0720/QcloudSms-go/QcloudSms"
	"gopkg.in/mgo.v2/bson"
)

const SMSExpireTime = 60

/**
发送SMS 存redis
*/
func SMSHandler(c *gin.Context) {
	var user_request SmsRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	// username唯一
	count, err := db.FindCount(db.DB, db.CollectionUser, bson.M{"tel": user_request.Tel})
	if err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "db err", Payload: nil})
		return
	}

	if count != 0 {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user already registered", Payload: nil})
		return
	}

	redis := utils.RedisUtils{}
	redis.Connect()
	defer redis.Close()
	app_id, _ := utils.GetConfig().GetInt("sms.app_id")
	app_key, _ := utils.GetConfig().Get("sms.app_key")

	var templID = 410376
	var params = []string{}
	qcloudsms, err := QcloudSms.NewQcloudSms(int(app_id), app_key)
	if err != nil {
		panic(err)
	}

	// add salt && sha256 password
	code := utils.GenRandomNumStr(time.Now().UnixNano(), 4)
	params = append(params, code)

	sendCh := make(chan error, 1)

	qcloudsms.SmsSingleSender.SendWithParam(86, user_request.Tel, templID, params, "", "", "", func(err error, resp *http.Response, resData string) {
		if err != nil {
			fmt.Println("err: ", err)
		} else {
			fmt.Println("response data: ", utils.Unicode2utf8(resData))
		}
		sendCh <- err
	})

	select {
	case err := <-sendCh:
		if err != nil {
			fmt.Print(err.Error())
			c.JSON(http.StatusInternalServerError, api.JSONReply{ErrorCode: 0, ErrorDescription: "server send sms error", Payload: nil})
			return
		}

		redis.SetEx(user_request.Tel, 60, code)

		c.JSON(http.StatusOK, api.JSONReply{
			ErrorCode:        0,
			ErrorDescription: "success",
			Payload: struct {
				Code string `json:"code"`
			}{
				code,
			}})
	}

}
