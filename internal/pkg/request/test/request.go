package request

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/dema501/randomjoke/internal/pkg/request"
	"github.com/pkg/errors"
)

// FakeSuperAgent has been build for unit tests purpose
type FakeSuperAgent struct {
	body io.Reader
}

// Constructor..
func New(b *strings.Reader) request.Maker {
	return &FakeSuperAgent{
		body: b,
	}
}

// Implement interface
func (r *FakeSuperAgent) Get(url string, result interface{}, args ...interface{}) error {
	if err := json.NewDecoder(r.body).Decode(result); err != nil {
		return errors.Wrap(err, "Can't Unmarshal response")
	}

	return nil
}
