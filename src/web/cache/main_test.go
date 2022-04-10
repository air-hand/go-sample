package cache

import (
	"context"
	"log"
	"os"
	"testing"
)

type cacheTester struct {
}

func (t *cacheTester) clearAll() {
	ctx := context.Background()

	config := NewCacheConnectConfigFromEnv()
	client := NewCacheClient(config)
	defer client.Close()

	err := client.FlushDB(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (t *cacheTester) setup() {
	t.clearAll()
}

func (t *cacheTester) teardown() {
}

func TestMain(m *testing.M) {
	tester := &cacheTester{}
	tester.setup()
	code := m.Run()
	tester.teardown()
	os.Exit(code)
}
