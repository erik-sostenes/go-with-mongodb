package repository

import "time"

// MongoDB contains the environment variables to configure the mongodb connection
type MongoDB struct {
	Dns,
	User,
	Password,
	DatabaseName string
	ConnectTimeout time.Duration
}
