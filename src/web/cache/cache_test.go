package cache

import (
	"context"
	"testing"
)

const (
	TEST_CACHE_DB = 15
)

func TestConnectToCache(t *testing.T) {
	ctx := context.Background()

	config := NewCacheConnectConfigFromEnv()
	config.DB = TEST_CACHE_DB
	client := NewCacheClient(config)
	defer client.Close()

	err := client.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		t.Error(err)
	}
}
