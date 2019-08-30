package user_api

// Create User
// we get uid from the jwt
type CreateUserRequest struct {
	//Tel      string `bson:"tel" json:"tel" binding:"required"`
	//Code     string `json:"code" binding:"required"`
	Username string `bson:"username" json:"username" binding:"required"`
	Password string `bson:"password" json:"password" binding:"required"`
}

type SmsRequest struct {
	Tel string `bson:"tel" json:"tel" binding:"required"`
}

type DeleteUserRequest struct {
	UID string `json:"uid"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SetKeyUserRequest struct {
	Username string        `json:"username"`
	Password string        `json:"password"`
	OkexKey  OkexKeyDetail `bson:"okex_key" json:"okex_key"`
	HuobiKey OkexKeyDetail `bson:"huobi_key" json:"huobi_key"`
}

type SetPushURLUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PushURL  string `json:"push_url"`
}
