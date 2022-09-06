package find

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/erik-sostenes/curso-mongo/internal/model"
	"github.com/erik-sostenes/curso-mongo/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFindNextAccount(t *testing.T) {
	tsc := map[string]struct {
		account          account
		expectedAccounts model.Accounts
		expectedError    error
	}{
		"Given a collection that exists, return a list of accounts": {
			account: NewAccount(
				repository.NewMDB(repository.Config).Collection("accounts"),
			),
		},
		"Given a collection that does not exist, return no account": {
			account: NewAccount(
				repository.NewMDB(repository.Config).Collection("some_collection"),
			),
			expectedError: mongo.ErrNoDocuments,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			accounts, err := ts.account.FindNextAccount(context.TODO())
			if !errors.Is(err, ts.expectedError) {
				t.Fatalf("expected error %v, got %v error", ts.expectedError, err)
				t.SkipNow()
			}

			if reflect.TypeOf(accounts) != reflect.TypeOf(ts.expectedAccounts) {
				t.Fatalf("expected %T, got %T", ts.expectedAccounts, accounts)
			}
		})
	}
}

func TestFindAllAccounts(t *testing.T) {
	tsc := map[string]struct {
		account          account
		expectedAccounts model.Accounts
		expectedError    error
	}{
		"Given a collection that exists, return a list of accounts": {
			account: NewAccount(
				repository.NewMDB(repository.Config).Collection("accounts"),
			),
		},
		"Given a collection that does not exist, return no account": {
			account: NewAccount(
				repository.NewMDB(repository.Config).Collection("some_collection"),
			),
			expectedError: mongo.ErrNoDocuments,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			accounts, err := ts.account.FindAllAccounts(context.TODO())
			if !errors.Is(err, ts.expectedError) {
				t.Fatalf("expected error %v, got %v error", ts.expectedError, err)
				t.SkipNow()
			}

			if reflect.TypeOf(accounts) != reflect.TypeOf(ts.expectedAccounts) {
				t.Fatalf("expected %T, got %T", ts.expectedAccounts, accounts)
			}
		})
	}
}

func BenchmarkFindNextAccount(b *testing.B) {
	account := NewAccount(repository.NewMDB(repository.Config).Collection("accounts"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		account.FindNextAccount(context.TODO())
	}
}

func BenchmarkFindAllAccounts(b *testing.B) {
	account := NewAccount(repository.NewMDB(repository.Config).Collection("accounts"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		account.FindNextAccount(context.TODO())
	}
}
