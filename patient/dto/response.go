package dto

import "fmt"

// BaseResponse provides base response for patients
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
	Message     string `json:"message"`
	OK          bool   `json:"ok"`
	ID          string `json:"patient_id"`
	PatientTime string `json:"patient_time"`
}

// String provides string repsentation
func (cr *CreateResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", cr.Message, cr.OK)
}

// UpdateResponse provides create response
type UpdateResponse struct {
	Message    string `json:"message"`
	OK         bool   `json:"ok"`
	ID         string `json:"patient_id"`
	UpdateTime string `json:"update_time"`
}

// String provides string repsentation
func (ur *UpdateResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", ur.Message, ur.OK)
}
