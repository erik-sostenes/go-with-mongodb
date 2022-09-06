package repository

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// MongoDB contains the environment variables to configure the mongodb connection
type MongoDB struct {
	Dns,
	User,
	Password,
	DatabaseName string
	ConnectTimeout time.Duration
}

var (
	timeout, _ = strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))
	// Config settings to a mongodb connection
	Config = MongoDB{
		Dns:            fmt.Sprintf("mongodb://127.0.0.1:%s", os.Getenv("NoSQL_PORT")),
		DatabaseName:   os.Getenv("NoSQL_DATABASE"),
		ConnectTimeout: time.Duration(timeout),
	}
)
