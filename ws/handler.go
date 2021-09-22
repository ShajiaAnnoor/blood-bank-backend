package ws

type Handler interface {
	Handle(Client, *RequestDTO)
}

type HandlerFunc func(Client, *RequestDTO)

func (hf HandlerFunc) Handle(c Client, data *RequestDTO) {
	hf(c, data)
}
