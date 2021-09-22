package ws

import "encoding/json"

type RequestDTO struct {
	Kind   Kind              `json:"kind"`
	Values map[string]string `json:"values"`
}

func (r *RequestDTO) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, r)
}

type ChatDTO struct {
	Event   string `json:"event"`
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

func (c *ChatDTO) ToBytes() ([]byte, error) {
	return json.Marshal(c)
}

type FriendRequestDTO struct {
	Event  string `json:"event"`
	Action string `json:"action"`
}

func (f *FriendRequestDTO) ToBytes() ([]byte, error) {
	return json.Marshal(f)
}
