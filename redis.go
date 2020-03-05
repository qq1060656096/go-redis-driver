package redisManager

import (
	"encoding/json"
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

func (m *ConnectionManager) Length() int {
	return len(m.connList)
}

func (m ConnectionManager) String() string {
	type tmpT struct {
		HasClient bool
		Options struct{
			Addr string
			Password string
			DB int
		}
	}
	list := make(map[string]tmpT)
	for k, v := range m.connList {
		tmp := tmpT{}
		if v.client != nil {
			tmp.HasClient = true
			tmp.Options.Addr = (*v).options.Addr
			tmp.Options.Password = (*v).options.Password
			tmp.Options.DB = (*v).options.DB

		}
		list[k] = tmp
	}
	bytes, _ := json.Marshal(list)
	return string(bytes)
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
