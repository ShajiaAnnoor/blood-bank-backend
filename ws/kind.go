package ws

type Kind string

func (k Kind) String() string {
	return string(k)
}

const (
	KindAuthentication Kind = "authentication"
)
