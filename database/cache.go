package database

import (
	"context"

	"github.com/Calgorr/URL_Shortener/model"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func AddLink(link *model.Link) error {
	if rdb.Get(ctx, link.Address) != nil {
		return nil
	}
	addLink(link)
	return rdb.Set(ctx, link.Hash, link.Address, 0).Err()
}

func GetLink(hash string) (string, error) {
	if val := rdb.Get(ctx, hash).Val(); val != "" {
		return val, nil
	}
	link, err := getLink(hash)
	if err != nil {
		return "", err
	}
	rdb.Set(ctx, hash, link.Address, 0)
	return link.Address, nil
}
