package go_redis_pool

import (
	"time"
	"errors"
)

var errWrongArguments error = errors.New("wrong number of arugments")

type Options struct {
	MaxIdle        int
	MaxActive      int
	Wait           bool
	IdleTimeout    time.Duration
	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func NewOptions() Options {
	return Options{
		MaxIdle:        64,
		MaxActive:      256,
		IdleTimeout:    180 * time.Second,
		ConnectTimeout: 1 * time.Second,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
	}
}

func FillOptions(opts Options) Options {
	initOpts := NewOptions()
	if opts.MaxIdle > 0 {
		initOpts.MaxIdle = opts.MaxIdle
	}
	if opts.MaxActive > 0 {
		initOpts.MaxActive = opts.MaxActive
	}
	if opts.Wait {
		initOpts.Wait = opts.Wait
	}
	if opts.IdleTimeout > 0 {
		initOpts.IdleTimeout = opts.IdleTimeout
	}
	if opts.ConnectTimeout > 0 {
		initOpts.ConnectTimeout = opts.ConnectTimeout
	}
	if opts.ReadTimeout > 0 {
		initOpts.ReadTimeout = opts.ReadTimeout
	}
	if opts.WriteTimeout > 0 {
		initOpts.WriteTimeout = opts.WriteTimeout
	}
	return initOpts
}
