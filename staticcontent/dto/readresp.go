package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading notice
type ReadResp struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(staticcontent *model.StaticContent) {
	r.UserID = staticcontent.UserID
	r.Text = staticcontent.Text
	r.ID = staticcontent.ID
}
