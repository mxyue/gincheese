package util

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"gopkg.in/gin-gonic/gin.v1"
	"time"
)

var cacheBox *cache.Cache

func init() {
	cacheBox = cache.New(10*time.Minute, 30*time.Second)
}

func CacheSet(namespace, key, value string) {
	cacheBox.Set(space_key(namespace, key), value, cache.DefaultExpiration)
	combine_key := space_key(namespace, key)
	code, _ := cacheBox.Get(combine_key)
	fmt.Fprintf(gin.DefaultWriter, "CacheSet get %s => value: %s \r\n", key, code)
}

func CacheGet(namespace, key string) (interface{}, bool) {
	return cacheBox.Get(space_key(namespace, key))
}

func space_key(namespace, key string) string {
	return fmt.Sprintf("%s-%s", namespace, key)
}

func CacheDelete(namespace, key string) {
	cacheBox.Delete(space_key(namespace, key))
}
