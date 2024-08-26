package migrations

import (
	"context"
	"github.com/golang-migrate/migrate/v4/source/golang"
	"github.com/redis/go-redis/v9"
)

func init() {
	golang.Register(Down1)
}

func Down1(ctx context.Context, client redis.UniversalClient) error {
	return client.Ping(ctx).Err()
}
