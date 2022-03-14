package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading patient
type ReadResp struct {
	Patient    string `json:"patient"`
	Sender     string `json:"sender"`
	Name       string `json:"name"`
	BloodGroup string `json:"blood_group"`
	District   string `json:"district"`
	Phone      string `json:"phone_number"`
	Address    string `json:"address"`
	UserID     string `json:"user_id"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(patient *model.Patient) {
	r.Patient = patient.Patient
	r.Sender = patient.UserID
}
