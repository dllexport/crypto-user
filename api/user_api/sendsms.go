package user_api

import (
	"fmt"
	"net/http"
	"time"

	"crypto-user/db"
	"crypto-user/utils"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const SMSExpireTime = 60

/**
发送SMS 存redis
*/
func SMSHandler(c *gin.Context) {
	var user_request SmsRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	// username唯一
	count, err := db.FindCount(db.DB, db.CollectionUser, bson.M{"tel": user_request.Tel})
	if err != nil {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "db err", Payload: nil})
		return
	}

	if count != 0 {
		c.JSON(http.StatusBadRequest, JSONReply{ErrorCode: -1, ErrorDescription: "user already registered", Payload: nil})
		return
	}

	redis := utils.RedisUtils{}
	redis.Connect()
	defer redis.Close()
	access_id, _ := utils.GetConfig().Get("sms.access_id")
	access_secret, _ := utils.GetConfig().Get("sms.access_secret")

	// add salt && sha256 password
	code := utils.GenRandomNumStr(time.Now().UnixNano(), 4)

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", access_id, access_secret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = user_request.Tel
	request.SignName = "共识区块链科技"
	request.TemplateCode = "SMS_173155142"
	codeSend := fmt.Sprintf("{\"code\":\"%s\"}", code)
	request.TemplateParam = codeSend
	fmt.Printf("%v\n", request)
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, JSONReply{ErrorCode: 0, ErrorDescription: "server send sms error", Payload: nil})
		return
	}

	fmt.Printf("response is %#v\n", response)
	if response.Code != "OK" {
		c.JSON(http.StatusInternalServerError, JSONReply{ErrorCode: 0, ErrorDescription: "fail to send sms", Payload: nil})
		return
	}

	redis.SetEx(user_request.Tel, SMSExpireTime, code)

	c.JSON(http.StatusOK, JSONReply{
		ErrorCode:        0,
		ErrorDescription: "success",
		Payload: struct {
			Code string `json:"code"`
		}{
			code,
		}})

}
