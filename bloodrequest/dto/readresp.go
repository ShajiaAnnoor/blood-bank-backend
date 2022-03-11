package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading blood request
type ReadResp struct {
	Request string `json:"request"`
	Sender  string `json:"sender"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(bloodreq *model.BloodRequest) {
	r.Request = bloodreq.Request
	r.Sender = bloodreq.UserID
}
