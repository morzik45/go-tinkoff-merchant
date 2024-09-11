package token_test

import (
	"testing"

	"github.com/go-pkgz/lgr"
	"github.com/morzik45/go-tinkoff-merchant/lib/utils/token"
)

func Test_New(t *testing.T) {
	tok, err := token.New(lgr.Default(), "https://sm-register-test.tscbank.ru/oauth/token", "aceplace", "SysANx")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("token: ", tok.Get())
}
