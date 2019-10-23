package user_api

import "time"

type OkexKeyDetail struct {
	APIKEY     string `bson:"api_key" json:"api_key"`
	SecretKey  string `bson:"secret_key" json:"secret_key"`
	PassPhrase string `bson:"passphrase" json:"passphrase"`
}
type HuobiKeyDetail struct {
	APIKEY    string `bson:"api_key" json:"api_key"`
	SecretKey string `bson:"secret_key" json:"secret_key"`
}

type User struct {
	UID           string         `bson:"_id" json:"uid"`
	Tel           string         `bson:"tel" json:"tel"`
	Password      string         `bson:"password" json:"password"` // hex
	Salt          string         `bson:"salt" json:"salt"`
	OkexKey       OkexKeyDetail  `bson:"okex_key" json:"okex_key"`
	HuobiKey      HuobiKeyDetail `bson:"huobi_key" json:"huobi_key"`
	PushURL       string         `bson:"push_url" json:"push_url"`
	Status        string         `bson:"status" json:"status"`
	CreatedTS     time.Time      `bson:"created_ts" json:"created_ts"`
	AllowStrategy []string       `bson:"allow_strategy" json:"allow_strategy"`
}

const (
	USER_STATUS_DELETED = "deleted"
	USER_STATUS_ACTIVE  = "active"
)
