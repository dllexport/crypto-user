package user_api

import "time"

type OkexKeyDetail struct {
	APIKEY     string `bson:"api_key" json:"api_key"`
	SecretKey  string `bson:"secret_key" json:"secret_key"`
	PassPhrase string `bson:"passphrase" json:"passphrasep;l.gh"`
}
type HuobiKeyDetail struct {
	APIKEY    string `bson:"api_key" json:"api_key"`
	SecretKey string `bson:"secret_key" json:"secret_key"`
}

type User struct {
	UID       string        `bson:"_id" json:"uid"`
	Username  string        `bson:"username" json:"username"`
	Password  string        `bson:"password" json:"password"` // hex
	Salt      string        `bson:"salt" json:"salt"`
	OkexKey   OkexKeyDetail `bson:"okex_key" json:"okex_key"`
	HuobiKey  OkexKeyDetail `bson:"huobi_key" json:"huobi_key"`
	Status    string        `bson:"status" json:"status"`
	CreatedTS time.Time     `bson:"created_ts" json:"created_ts"`
}

const (
	USER_STATUS_DELETED = "deleted"
	USER_STATUS_ACTIVE  = "active"
)
