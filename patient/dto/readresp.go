package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading patient
type ReadResp struct {
	Patient string `json:"patient"`
	Sender  string `json:"sender"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(patient *model.Patient) {
	r.Patient = patient.Patient
	r.Sender = patient.UserID
}
