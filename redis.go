package redis_driver

import (
	"github.com/go-redis/redis"
)

type Driver struct {
	drivers map[string]*Connection
}

type Connection struct {
	client  *redis.Client
	options *redis.Options
}

func NewDriver() *Driver {
	d := Driver{
		drivers: make(map[string]*Connection),
	}
	return &d
}

func (d *Driver) Add(name string, options *redis.Options) {
	d.drivers[name] = &Connection{
		options: options,
	}
}

func (d *Driver) Remove(name string) {
	delete(d.drivers, name)
}

func (d *Driver) Get(name string) *Connection {
	con, ok := d.drivers[name]
	if !ok {
		return nil
	}
	return con
}

func (d *Driver) Exist(name string) bool {
	con := d.Get(name)
	if con == nil {
		return false
	}
	return true
}

func (d *Driver) length() int {
	return len(d.drivers)
}

func (d *Driver) Reconnection(name string) *Connection {
	con := d.Get(name)
	if con == nil {
		return con
	}
	con.client = redis.NewClient(con.options)
	return con
}

func (c *Connection) GetRedisClient() *redis.Client {
	if c.client == nil {
		c.ReconnectRedisClient()
	}
	return c.client
}

func (c *Connection) ReconnectRedisClient() {
	c.client = redis.NewClient(c.options)
}

func (c *Connection) DisconnectRedisClient() bool {
	if c.client != nil {
		c.client.Close()
		c.client = nil
	}
	return true
}
