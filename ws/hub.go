package ws

type Hub interface {
	Send(userID string, data []byte) error
	HandleClient(c Client)
}
