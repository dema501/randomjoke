package jokemaker

import (
	"strings"
	"testing"

	"github.com/dema501/randomjoke/internal/pkg/request"
)

func TestRandomJokeGenerator(t *testing.T) {
	payload := `{ "type": "success", "value": { "id": 534, "joke": "John Doe is the ultimate mutex, all threads fear him.", "categories": ["nerdy"] } }`

	sa := &request.FakeSuperAgent{
		Body: strings.NewReader(payload),
	}

	rj := New(sa)

	if err := rj.Generate(); err != nil {
		t.Errorf("Expected NoError %v", err)
	}

	joke := rj.GetJoke()
	if strings.Contains(joke, "ultimate mutex") == false {
		t.Errorf("Expected Joke ...ultimate mutex... but get %v", joke)
	}
}
