package redis

import(
	"github.com/redis/go-redis/v9"

	"github.com/valikhan03/auctions-service/models"
)

func InitRedisConn() *redis.Client {
	opts := new(redis.Options)
	opts.Network = models.RedisConfigGlobal.Network
	opts.Addr = models.RedisConfigGlobal.Addr
	opts.ClientName = models.RedisConfigGlobal.ClientName
	opts.Username = models.RedisConfigGlobal.Username
	opts.DB = models.RedisConfigGlobal.DB

	client := redis.NewClient(opts)

	return client
}
