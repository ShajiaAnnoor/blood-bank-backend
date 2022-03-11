package dto

import "fmt"

// CreateResponse provides create response
type DeleteResponse struct {
	Message     string `json:"message"`
	OK          bool   `json:"ok"`
	ID          string `json:"blood_request_id"`
	RequestTime string `json:"request_time"`
}

// String provides string repsentation
func (c *DeleteResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", c.Message, c.OK)
}
