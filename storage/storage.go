package storage

import (
	"time"
	"github.com/dbv/dspider/modinit"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func CacheSet(key string, data interface{}, duration time.Duration) {
	for {
		if rediscmd := modinit.RedisInstance.Set(key, data, duration); rediscmd.Err() != nil {
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}
}

func CacheGet(key string, data *[]byte) {
	var e error
	for {
		if *data, e = modinit.RedisInstance.Get(key).Bytes(); e != nil {
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}
}
