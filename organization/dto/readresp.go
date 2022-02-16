package dto

import "gitlab.com/Aubichol/hrishi-backend/model"

//ReadResp holds the response data for reading organization
type ReadResp struct {
	Organization string `json:"organization"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(org *model.Comment) {
	r.Organization = org.Organization
}
