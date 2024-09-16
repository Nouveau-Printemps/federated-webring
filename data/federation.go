package data

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

const (
	FederationRequestType  = "federation-request"
	FederationResponseType = "federation-response"
	FederationBreakType    = "federation-break"
)

var ErrNotFound = errors.New("federation sent not found")

func GetFederationSent(typeSent string, origine string) (string, error) {
	val := rdb.Get(context.Background(), "federation:"+typeSent+":"+origine)
	if val.Err() != nil {
		if errors.Is(val.Err(), redis.Nil) {
			return "", ErrNotFound
		}
		return "", val.Err()
	}
	return val.Val(), nil
}
