package request

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// FakeSuperAgent has been build for unit tests purpose
type FakeSuperAgent struct {
	Body io.Reader
}

// Implement interface
func (r *FakeSuperAgent) Get(url string, result interface{}, args ...interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(result); err != nil {
		return errors.Wrap(err, "Can't Unmarshal response")
	}

	return nil
}
