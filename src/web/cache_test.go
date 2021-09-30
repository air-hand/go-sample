package web

import "testing"

const (
	TEST_CACHE_DB = 15
)

func TestConnectToCache(t *testing.T) {
	config := NewCacheConnectConfigFromEnv()
	config.DB = TEST_CACHE_DB
	conn := NewCacheClient(config)
	defer conn.Close()
}
