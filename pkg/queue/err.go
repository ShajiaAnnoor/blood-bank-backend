package queue

type Error string

func (s Error) Error() string {
	return string(s)
}

// list of error
const (
	Closed = Error("closed")
)
