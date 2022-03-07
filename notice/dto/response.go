package dto

import "fmt"

// BaseResponse provides base response for notices
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
	Message    string `json:"message"`
	OK         bool   `json:"ok"`
	ID         string `json:"notice_id"`
	NoticeTime string `json:"notice_time"`
}

// String provides string repsentation
func (cr *CreateResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", cr.Message, cr.OK)
}

// UpdateResponse provides create response
type UpdateResponse struct {
	Message    string `json:"message"`
	OK         bool   `json:"ok"`
	ID         string `json:"notice_id"`
	UpdateTime string `json:"update_time"`
}

// String provides string repsentation
func (ur *UpdateResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", ur.Message, ur.OK)
}
