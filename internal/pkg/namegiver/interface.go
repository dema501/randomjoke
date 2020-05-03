package namegiver

type Doer interface {
	GetName() (string, string)
	Generate() error
}