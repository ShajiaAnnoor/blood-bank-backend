package donor

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/donor/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storedonor "gitlab.com/Aubichol/blood-bank-backend/store/donor"
	"go.uber.org/dig"
)

//Reader provides an interface for reading donores
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//donorReader implements Reader interface
type donorReader struct {
	donors storedonor.Donors
}

func (read *donorReader) askStore(donorID string) (
	donor *model.Donor,
	err error,
) {
	donor, err = read.donors.FindByID(donorID)
	return
}

func (read *donorReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *donorReader) prepareResponse(
	donor *model.Donor,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(donor)
	return
}

func (read *donorReader) Read(donorReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	donor, err := read.askStore(donorReq.DonorID)
	if err != nil {
		logrus.Error("Could not find donor error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(donor)

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Donor storedonor.Donors
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &donorReader{
		donors: params.Donor,
	}
}
