package db

import "crypto-user/utils"

var (
	DB             = "tradedb"
	CollectionUser = "user"
)

func init() {
	DB, _ = utils.GetConfig().Get("db.name")
}
