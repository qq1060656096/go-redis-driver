package manager

import (
	"github.com/go-redis/redis"
)

type ConnectionManager struct {
	connList map[string]*Connection
}

type Connection struct {
	client  *redis.Client
	options *redis.Options
}

func NewConnectionManager() *ConnectionManager {
	m := ConnectionManager{
		connList: make(map[string]*Connection),
	}
	return &m
}

func (m *ConnectionManager) Add(name string, options *redis.Options) {
	m.connList[name] = &Connection{
		options: options,
	}
}

func (m *ConnectionManager) Remove(name string) {
	delete(m.connList, name)
}

func (m *ConnectionManager) Get(name string) *Connection {
	con, ok := m.connList[name]
	if !ok {
		return nil
	}
	return con
}

func (m *ConnectionManager) Exist(name string) bool {
	con := m.Get(name)
	if con == nil {
		return false
	}
	return true
}

func (m *ConnectionManager) length() int {
	return len(m.connList)
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
