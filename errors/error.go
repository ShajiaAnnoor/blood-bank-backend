package errors

import "fmt"

type Base struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

func (b *Base) String() string {
	return fmt.Sprintf("message:%s, ok:%v", b.Message, b.OK)
}

type Invalid struct {
	Base
}

func (i *Invalid) Error() string {
	return "invalid : " + i.String()
}

type Unknown struct {
	Base
}

func (u *Unknown) Error() string {
	return "unknown : " + u.String()
}

type Unauthorized struct {
	Base
}

func (u *Unauthorized) Error() string {
	return "unauthorized : " + u.String()
}
