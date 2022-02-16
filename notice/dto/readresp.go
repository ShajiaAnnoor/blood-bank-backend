package dto

import "gitlab.com/Aubichol/hrishi-backend/model"

//ReadResp holds the response data for reading notice
type ReadResp struct {
	Notice string `json:"notice"`
	Sender string `json:"sender"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(notice *model.Notice) {
	r.Notice = notice.Notice
	r.Sender = notice.UserID
}
