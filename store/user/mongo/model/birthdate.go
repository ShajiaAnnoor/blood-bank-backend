package model

import (
	"gitlab.com/Aubichol/hrishi-backend/model"
)

//BirthDate defines mongodb data type for BirthDate
type BirthDate struct {
	Year  int `bson:"year"`
	Month int `bson:"month"`
	Day   int `bson:"day"`
}

//FromModel converts model data to mongodb model data for a user's date of birth
func (b *BirthDate) FromModel(data *model.BirthDate) {
	b.Year = data.Year
	b.Month = data.Month
	b.Day = data.Day
}

//ModelBirthDate converts bson to model for birth date
func (b *BirthDate) ModelBirthDate() *model.BirthDate {
	data := model.BirthDate{}
	data.Year = b.Year
	data.Month = b.Month
	data.Day = b.Day
	return &data
}
