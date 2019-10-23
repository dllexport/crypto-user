package user_api

// Create User
// we get uid from the jwt
type CreateUserRequest struct {
	Tel      string `bson:"tel" json:"tel" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

type SmsRequest struct {
	Tel string `bson:"tel" json:"tel" binding:"required"`
}

type DeleteUserRequest struct {
	UID string `json:"uid"`
}

type LoginUserRequest struct {
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type SetKeyUserRequest struct {
	OkexKey  OkexKeyDetail  `bson:"okex_key" json:"okex_key"`
	HuobiKey HuobiKeyDetail `bson:"huobi_key" json:"huobi_key"`
	PushURL  string `json:"push_url"`
}

type SetPushURLUserRequest struct {
	Tel      string `json:"tel"`
	Password string `json:"password"`
	PushURL  string `json:"push_url"`
}
