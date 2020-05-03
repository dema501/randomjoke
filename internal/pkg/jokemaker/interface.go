package jokemaker

type Doer interface {
	GetJoke() string
	SetName(string, string)
	Generate() error 
}