package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading notice
type ReadResp struct {
	PatientName string `json:"patient_name"`
	BloodGroup  string `json:"notice_group"`
	Description string `json:"description"`
	District    string `json:"district"`
	Address     string `json:"address"`
	Title       string `json:"title"`
	UserID      string `json:"user_id"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(notice *model.Notice) {
	//r.Notice = notice.Notice
	//r.Sender = notice.UserID
}
