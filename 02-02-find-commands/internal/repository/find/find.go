package find

import (
	"context"

	"github.com/erik-sostenes/curso-mongo/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type account struct {
	*mongo.Collection
}

func NewAccount(mongo *mongo.Collection) account {
	return account{mongo}
}

func (a *account) FindNextAccount(ctx context.Context) (accounts model.Accounts, err error) {
	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := a.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		var account model.Account
		if err = cur.Decode(&account); err != nil {
			return
		}
		accounts = append(accounts, account)
	}

	if err = cur.Err(); err != nil {
		return
	}

	if len(accounts) == 0 {
		err = mongo.ErrNoDocuments
		return
	}
	return
}

func (a *account) FindAllAccounts(ctx context.Context) (accounts model.Accounts, err error) {
	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := a.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	// If your expected result set is large, using the *mongo.Cursor.All function might not be the best idea
	if err = cur.All(ctx, &accounts); err != nil {
		return
	}

	if len(accounts) == 0 {
		err = mongo.ErrNoDocuments
	}
	return
}
