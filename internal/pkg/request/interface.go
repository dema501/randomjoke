package request

type Maker interface {
	Get(url string, result interface{}, args ...interface{}) error
}
