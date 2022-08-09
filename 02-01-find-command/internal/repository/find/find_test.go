package find

import (
	"context"
	"errors"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/erik-sostenes/curso-mongo/internal/model"
	"github.com/erik-sostenes/curso-mongo/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFindAllAccounts(t *testing.T) {
	timeout, _ := strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))

	config := repository.MongoDB{
		Uri:            "mongodb+srv://%s:%s@cluster0.o3tc8ee.mongodb.net/?retryWrites=true&w=majority",
		User:           os.Getenv("NoSQL_USER"),
		Password:       os.Getenv("NoSQL_PASSWORD"),
		DatabaseName:   os.Getenv("NoSQL_DATABASE"),
		ConnectTimeout: time.Duration(timeout),
	}
	tsc := map[string]struct {
		accounts         Account
		expectedAccounts model.Accounts
		expectedError    error
	}{
		"Given a collection that exists, return a list of accounts": {
			accounts: NewAccount(
				repository.NewMDB(config).Collection("accounts"),
			),
		},
		"Given a collection that does not exist, return no account": {
			accounts: NewAccount(
				repository.NewMDB(config).Collection("some_collection"),
			),
			expectedError: mongo.ErrNoDocuments,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			accounts, err := ts.accounts.FindAllAccounts(context.TODO())
			if err != nil {
				if !errors.Is(err, ts.expectedError) {
					t.Fatalf("expected error %v, got %v error", ts.expectedError, err)
				}
				t.SkipNow()
			}

			if reflect.TypeOf(accounts) != reflect.TypeOf(ts.expectedAccounts) {
				t.Fatalf("expected %T, got %T", ts.expectedAccounts, ts.accounts)
			}
		})
	}
}
