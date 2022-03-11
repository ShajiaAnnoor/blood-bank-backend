package dto

import "gitlab.com/Aubichol/blood-bank-backend/model"

//BirthDate stores birth data data
type BirthDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

//FromModel converts model data to json type data
func (b *BirthDate) FromModel(data *model.BirthDate) {
	b.Year = data.Year
	b.Month = data.Month
	b.Day = data.Day
}

//ToModel converts json type data to model data
func (b *BirthDate) ToModel(m *model.BirthDate) {
	m.Year = b.Year
	m.Month = b.Month
	m.Day = b.Day
}
