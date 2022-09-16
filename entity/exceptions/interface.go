package exceptions

type Error interface {
	StatusCode() int
}
