package redis_manager

import (
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	d := NewConnectionManager()
	d.Add("test1", &redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       1,        // use default DB
	})
	d.Add("test2", &redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       2,        // use default DB
	})
	assert.Equal(t, 2, d.length(), "driver.length.func.error")

	con := d.Get("test3.driverNotExist")
	isNil := false
	if con == nil {
		isNil = true
	}
	assert.Equal(t, true, isNil, "driver.get.func.test3.driverNotExist")

	con = d.Get("test2")
	isNil = false
	if con == nil {
		isNil = true
	}
	assert.Equal(t, false, isNil, "driver.get.func.test2.driverExist")

	err := d.Get("test1").GetRedisClient().Set("test1.key1", "test1.value1", 0).Err()
	assert.Equal(t, nil, err, "driver.get.test1.setRedisKeyValue.error")
	d.Get("test2").DisconnectRedisClient()
	err = d.Get("test2").GetRedisClient().Set("test2.key2", "test2.value2.1", 0).Err()
	assert.Equal(t, nil, err, "driver.get.test2.setRedisKeyValue.error")
}
