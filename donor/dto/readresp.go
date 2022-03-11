package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading donor
type ReadResp struct {
	Donor  string `json:"donor"`
	Sender string `json:"sender"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(donor *model.Donor) {
	//	r.Donor = donor.Donor
	r.Sender = donor.UserID
}
