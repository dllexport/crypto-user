package user_api

// Create User
// we get uid from the jwt
type CreateUserRequest struct {
	Username string        `bson:"username" json:"username" binding:"required"`
	Password string        `bson:"password" json:"password" binding:"required"`
	OkexKey  OkexKeyDetail `bson:"okex_key" json:"okex_key"`
	HuobiKey OkexKeyDetail `bson:"huobi_key" json:"huobi_key"`
}

type DeleteUserRequest struct {
	UID string `json:"uid"`
}
