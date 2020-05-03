package jokemaker

import (
	"strings"
	"testing"

	request "github.com/dema501/randomjoke/internal/pkg/request/test"
)

func TestRandomJokeGenerator(t *testing.T) {
	payload := `{ "type": "success", "value": { "id": 534, "joke": "John Doe is the ultimate mutex, all threads fear him.", "categories": ["nerdy"] } }`

	sa := request.New(strings.NewReader(payload))

	rj := New(sa)

	if err := rj.Generate(); err != nil {
		t.Errorf("Expected NoError %v", err)
	}

	joke := rj.GetJoke()
	if strings.Contains(joke, "ultimate mutex") == false {
		t.Errorf("Expected Joke ...ultimate mutex... but get %v", joke)
	}
}
