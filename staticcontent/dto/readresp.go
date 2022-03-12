package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading notice
type ReadResp struct {
	Notice string `json:"staticcontent"`
	Sender string `json:"sender"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(staticcontent *model.StaticContent) {
	//	r.Notice = staticcontent.Notice
	r.Sender = staticcontent.UserID
}
