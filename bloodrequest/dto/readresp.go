package dto

import "gitlab.com/Aubichol/hrishi-backend/model"

//ReadResp holds the response data for reading comment
type ReadResp struct {
	Request string `json:"request"`
	Sender  string `json:"sender"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(comment *model.Comment) {
	r.Request = comment.Request
	r.Sender = comment.UserID
}
