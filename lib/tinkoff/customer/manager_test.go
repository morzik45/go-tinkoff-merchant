package customer_test

import (
	"testing"

	"github.com/go-pkgz/lgr"
	"github.com/morzik45/go-tinkoff-merchant/lib/tinkoff"
)

const (
	testTerminalKey = "1683019138816DEMO"
	testPassword    = "83qy5pk7qzosss0m"
)

func TestManager_AddCardLink(t *testing.T) {
	c := tinkoff.New(false, lgr.Default()).Customer(
		testTerminalKey,
		testPassword,
	)

	resp, err := c.AddCard("reansnow")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	t.Log(resp)
}
