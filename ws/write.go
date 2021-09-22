package ws

type Client interface {
	Write([]byte) error
	Read() ([]byte, error)
	Kick() error

	ID() string
	SetID(string)
}
