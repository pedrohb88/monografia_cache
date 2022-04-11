package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/joho/godotenv"
)

type Cache struct {
	client *memcache.Client
}

func New() (*Cache, error) {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading config; %w", err)
		}
	}

	mc := memcache.New(os.Getenv("CACHECONN"))
	return &Cache{
		client: mc,
	}, nil
}

func (c *Cache) GetSet(key string, obj interface{}, getDataFunc func() (interface{}, error)) error {
	var shouldSetCache bool
	var o interface{}
	err := c.Get(key, obj)
	if errors.Is(err, memcache.ErrCacheMiss) {
		var e error

		o, e = getDataFunc()

		if e != nil {
			return e
		}
		shouldSetCache = true
	} else if err != nil {
		return err
	}

	if shouldSetCache {
		err = c.Set(key, o)
		if err != nil {
			return err
		}
		return c.Get(key, obj)
	}

	return nil
}

func (c *Cache) GetMulti(keys []string, populateFn func(key string) interface{}) error {

	objMap, err := c.client.GetMulti(keys)
	if err != nil {
		return err
	}

	for key, value := range objMap {
		obj := populateFn(key)
		err = json.Unmarshal(value.Value, obj)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Cache) Get(key string, obj interface{}) error {
	item, err := c.client.Get(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(item.Value, obj)
}

func (c *Cache) Set(key string, obj interface{}) error {

	v, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return c.client.Set(&memcache.Item{Key: key, Value: v})
}

func (c *Cache) Delete(key string) error {
	err := c.client.Delete(key)
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return err
	}
	return nil
}

func Key(key string, digest string, args ...interface{}) string {
	fullKey := fmt.Sprintf("%s-%s", key, digest)
	for _, v := range args {
		fullKey += fmt.Sprintf("-%v", v)
	}
	return fullKey
}

func Keys(key string, digest string, args ...interface{}) []string {
	keys := make([]string, len(args))
	for i, v := range args {
		keys[i] = fmt.Sprintf("%s-%s-%v", key, digest, v)
	}
	return keys
}
