package namegiver

import (
	"fmt"

	"github.com/dema501/randomjoke/internal/pkg/request"
)

type randomName struct {
	SuperAgent request.Maker
	url        string
	response   struct {
		Results []struct {
			Name struct {
				First string `json:"first"`
				Last  string `json:"last"`
			} `json:"name"`
		} `json:"results"`
	}
}

func New(sa request.Maker) Doer {
	r := randomName{}
	r.SuperAgent = sa
	r.url = "https://randomuser.me/api/"
	return &r
}

func (rn *randomName) GetName() (string, string) {
	return rn.response.Results[0].Name.First, rn.response.Results[0].Name.Last
}

func (rn *randomName) Generate() error {
	if err := rn.SuperAgent.Get(rn.url, &rn.response); err != nil {
		return err
	}

	if len(rn.response.Results) == 0 {
		return fmt.Errorf("Results in random name response equal zero: %v", rn.response)
	}

	return nil
}
