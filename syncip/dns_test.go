package syncip_test

import (
	"context"
	"github.com/iikira/ippush/syncip"
	"log"
	"os"
	"testing"
)

var (
	apiKey   = os.Getenv("CF_API_KEY")
	apiEmail = os.Getenv("CF_API_EMAIL")
	ctx      = context.Background()
)

func TestTxt(t *testing.T) {
	sync, err := syncip.NewSyncIP(apiKey, apiEmail)
	if err != nil {
		log.Fatalln(err)
	}

	err = sync.SetZone(ctx, "iikira.com")
	if err != nil {
		log.Fatalln(err)
	}

	err = sync.SetTXTRecord(ctx, "acme", "kasfasdfjlskdjflsdjf")
	if err != nil {
		log.Fatalln(err)
	}
}
