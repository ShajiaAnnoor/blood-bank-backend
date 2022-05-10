package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//ReadResp holds the response data for reading donor
type ReadResp struct {
	Name         string `json:"name"`
	Phone        string `json:"phone_number"`
	District     string `json:"district"`
	BloodGroup   string `json:"blood_group"`
	Address      string `json:"address"`
	Availability bool   `json:"availability"`
	TimesDonated int    `json:"times_donated"`
	UserID       string `json:"user_id"`
}

//FromModel converts the model data to response data
func (r *ReadResp) FromModel(donor *model.Donor) {
	r.Name = donor.Name
	r.Phone = donor.Phone
	r.District = donor.District
	r.BloodGroup = donor.BloodGroup
	r.Address = donor.Address
	r.Availability = donor.Availability
	r.TimesDonated = donor.TimesDonated
	r.UserID = donor.UserID
}
