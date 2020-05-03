package jokemaker

import (
	"fmt"

	"github.com/dema501/randomjoke/internal/pkg/request"
	"github.com/pkg/errors"
)

type randomJoke struct {
	SuperAgent request.Maker
	url        string
	response   struct {
		Type  string `json:"type"`
		Value struct {
			ID   int    `json:"id"`
			Joke string `json:"joke"`
			// Categories []string `json:"categories"`
		} `json:"value"`
	}
}

func New(sa request.Maker) Doer {
	r := &randomJoke{}
	r.SuperAgent = sa
	r.url = "http://api.icndb.com/jokes/random?limitTo=[nerdy]"

	return r
}

func (rj *randomJoke) GetJoke() string {
	return rj.response.Value.Joke
}

func (rj *randomJoke) SetName(firstName, lastName string) {
	rj.url = fmt.Sprintf(
		"http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=[nerdy]",
		firstName,
		lastName,
	)
}

func (rj *randomJoke) Generate() error {
	if err := rj.SuperAgent.Get(rj.url, &rj.response); err != nil {
		return err
	}

	if rj.response.Type != "success" || rj.response.Value.Joke == "" {
		return errors.New("Results in random joke response is empty")
	}

	return nil
}
