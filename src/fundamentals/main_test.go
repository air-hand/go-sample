package fundamentals

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("setup")
	code := m.Run()
	log.Print("teardown")
	os.Exit(code)
}
