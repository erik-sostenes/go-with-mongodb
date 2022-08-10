package repository

import (
	"os"
	"strconv"
	"time"
)

// MongoDB contains the environment variables to configure the mongodb connection
type MongoDB struct {
	Uri,
	User,
	Password,
	DatabaseName string
	ConnectTimeout time.Duration
}

var (
	timeout, _ = strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))
	// Config settings to a mongodb connection
	Config = MongoDB{
		Uri:            "mongodb+srv://%s:%s@cluster0.o3tc8ee.mongodb.net/?retryWrites=true&w=majority",
		User:           os.Getenv("NoSQL_USER"),
		Password:       os.Getenv("NoSQL_PASSWORD"),
		DatabaseName:   os.Getenv("NoSQL_DATABASE"),
		ConnectTimeout: time.Duration(timeout),
	}
)
