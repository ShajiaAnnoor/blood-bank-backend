package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading organization
type ReadResp struct {
	Organization string `json:"organization"`
	Name         string `json:"name"`
	Phone        string `json:"phone_number"`
	District     string `json:"district"`
	Description  string `json:"description"`
	Address      string `json:"address"`
	UserID       string `json:"user_id"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(org *model.Organization) {
	r.Organization = org.Organization
}
