package repository

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestNewMongoClient(t *testing.T) {
	const dns = "mongodb://127.0.0.1:%s"
	timeout, _ := strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))

	tsc := map[string]struct {
		configuration MongoDB
		expectedError error
	}{
		"Given a correct configuration, a new mongo local client will be created": {
			configuration: MongoDB{
				Dns:            fmt.Sprintf(dns, os.Getenv("NoSQL_PORT")),
				DatabaseName:   os.Getenv("NoSQL_DATABASE"),
				ConnectTimeout: time.Duration(timeout),
			},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			_, err := NewMongoClient(ts.configuration)
			if err != ts.expectedError {
				t.Fatalf("expected error %T, got %T error", ts.expectedError, err)
			}
		})
	}
}
