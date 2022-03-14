package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading blood request
type ReadResp struct {
	Request    string `json:"request"`
	BloodGroup string `json:"blood_group"`
	UserID     string `json:"user_id"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(bloodreq *model.BloodRequest) {
	r.Request = bloodreq.Request
	//	r.Sender = bloodreq.UserID
}
