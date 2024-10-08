package tinkoff

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/morzik45/go-tinkoff-merchant/lib/tinkoff/customer"
	"github.com/morzik45/go-tinkoff-merchant/lib/tinkoff/notifications"
	"github.com/morzik45/go-tinkoff-merchant/lib/tinkoff/partner"
	"github.com/morzik45/go-tinkoff-merchant/lib/tinkoff/payments"
	"github.com/morzik45/go-tinkoff-merchant/lib/tinkoff/safedeal"
	"github.com/morzik45/go-tinkoff-merchant/lib/utils"
)

func New(debug bool, log utils.Logger) *Tinkoff {
	return &Tinkoff{
		debug: debug,
		log:   log,
		cl: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

type Tinkoff struct {
	debug bool
	log   utils.Logger
	cl    *http.Client
}

func (t *Tinkoff) Partner(username, password string) (*partner.Partner, error) {
	return partner.New(t, username, password)
}

func (t *Tinkoff) Payments(terminalID, password string) (*payments.Manager, error) {
	return payments.New(t, terminalID, password)
}

func (t *Tinkoff) Notifications(terminalID, password string) *notifications.Manager {
	return notifications.New(t.log, terminalID, password)
}

func (t *Tinkoff) Customer(terminalID, password string) *customer.Manager {
	return customer.New(t, terminalID, password)
}

func (t *Tinkoff) SafeDeal(terminalID, password string) (*safedeal.Manager, error) {
	return safedeal.New(t, terminalID, password)
}

func (t *Tinkoff) Debug() bool {
	return t.debug
}

func (t *Tinkoff) Log() utils.Logger {
	return t.log
}

func (t *Tinkoff) Request(r *http.Request, data any) error {
	r.Header.Add("Content-Type", "application/json")

	{
		b, _ := httputil.DumpRequest(r, true)
		t.log.Logf("[DEBUG] Request: %s", string(b))
	}

	resp, err := t.cl.Do(r)
	if err != nil {
		return err
	}

	{
		b, _ := httputil.DumpResponse(resp, true)
		t.log.Logf("[DEBUG] Response: %s", string(b))
	}

	decoder := json.NewDecoder(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		err := decoder.Decode(data)
		return err
	default:
		errorStruct := &ErrorResponse{}
		err := decoder.Decode(errorStruct)
		if err != nil {
			return err
		}

		return errorStruct
	}
}
