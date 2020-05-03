package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// SuperAgent it's wrapper around stdlib
// combines together few things http SuperAgent + unmarshalling response + basic validation
type SuperAgent struct {
	Header *http.Header
	Client *http.Client
}

type Header map[string]string

// {username,password}
type Auth []string

// Constructor..
func New() Maker {
	return &SuperAgent{
		Client: &http.Client{Timeout: time.Millisecond * 1200},
	}
}

func (r *SuperAgent) Get(url string, result interface{}, args ...interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "Can't make random name SuperAgent")
	}

	for _, arg := range args {
		switch a := arg.(type) {
		// arg is Header , set to SuperAgent header
		case Header:
			for k, v := range a {
				req.Header.Set(k, v)
			}
		case Auth:
			// a{username,password}
			req.SetBasicAuth(a[0], a[1])
		}
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "Can't make SuperAgent")
	}

	// yes sometimes resp.Body.Close() can have an error
	// it's better to log it, now it's just into stdout
	// TODO: pass logger into struct
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("[WARN] can't close response body, %s", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "Response status is not OK response")
	}

	// nothing to Unmarshal into
	if result == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return errors.Wrap(err, "Can't Unmarshal response")
	}

	return nil
}
