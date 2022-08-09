package repository

import (
	"errors"
	"os"
	"strconv"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

func TestNewMongoClient(t *testing.T) {
	const uri = "mongodb+srv://%s:%s@cluster0.o3tc8ee.mongodb.net/?retryWrites=true&w=majority"
	timeout, _ := strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))

	tsc := map[string]struct {
		configuration MongoDB
		// expected is the type of error to expect
		expectedError topology.ConnectionError
	}{
		"Given a correct configuration, a new mongo client will be created": {
			configuration: MongoDB{
				Uri:            uri,
				User:           os.Getenv("NoSQL_USER"),
				Password:       os.Getenv("NoSQL_PASSWORD"),
				DatabaseName:   os.Getenv("NoSQL_DATABASE"),
				ConnectTimeout: time.Duration(timeout),
			},
		},
		"Given incorrect authentication configuation, a new mongo client will not be created": {
			configuration: MongoDB{
				Uri:            uri,
				User:           "some_user",
				Password:       "some_user",
				DatabaseName:   os.Getenv("NoSQL_DATABASE"),
				ConnectTimeout: time.Duration(timeout),
			},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			_, err := NewMongoClient(ts.configuration)
			if err != nil {
				if !errors.As(err, &ts.expectedError) {
					t.Fatalf("expected error %T, got %T error", ts.expectedError, err)
				}
			}
		})
	}
}
