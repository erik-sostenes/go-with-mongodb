package model

type (
	// Account is an object model.
	Account struct {
		AccountId int      `bson:"account_id"`
		Limit     int      `bson:"limit"`
		Products  []string `bson:"products"`
	}
	// Accounts represents a slice of Account
	Accounts []Account
)
