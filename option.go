package go_redis_pool

import (
	"time"
	"errors"
)

var errWrongArguments error = errors.New("wrong number of arugments")

type Options struct {
	DB             int
	MaxIdle        int
	MaxActive      int
	Wait           bool
	IdleTimeout    time.Duration
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	Password       string
}

func NewOptions() *Options {
	return &Options{
		DB:             0,
		MaxIdle:        256,
		MaxActive:      1024,
		IdleTimeout:    180 * time.Second,
		ConnectTimeout: 1 * time.Second,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		Password:       "",
	}
}
