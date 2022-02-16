package dto

import "fmt"

// BaseResponse provides base response for donors
type BaseResponse struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

// String provides string repsentation
func (b *BaseResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", b.Message, b.OK)
}

// CreateResponse provides create response
type CreateResponse struct {
	Message   string `json:"message"`
	OK        bool   `json:"ok"`
	ID        string `json:"donor_id"`
	DonorTime string `json:"donor_time"`
}

// String provides string repsentation
func (c *CreateResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", c.Message, c.OK)
}

// UpdateResponse provides create response
type UpdateResponse struct {
	Message    string `json:"message"`
	OK         bool   `json:"ok"`
	ID         string `json:"donor_id"`
	UpdateTime string `json:"update_time"`
}

// String provides string repsentation
func (c *UpdateResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", c.Message, c.OK)
}
