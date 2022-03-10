package bloodrequest

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storebloodrequest "gitlab.com/Aubichol/blood-bank-backend/store/bloodrequest"
	"go.uber.org/dig"
)

//Reader provides an interface for reading bloodrequestes
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//bloodrequestReader implements Reader interface
type bloodrequestReader struct {
	bloodrequestes bloodrequest.BloodRequest
}

func (read *bloodrequestReader) askStore(bloodrequestID string) (
	bloodrequest *model.BloodRequest,
	err error,
) {
	bloodrequest, err = read.bloodrequestes.FindByID(bloodrequestID)
	return
}

func (read *bloodrequestReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *bloodrequestReader) prepareResponse(
	bloodrequest *model.BloodRequest,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(bloodrequest)
	return
}

func (read *bloodrequestReader) Read(bloodrequestReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	bloodrequest, err := read.askStore(bloodrequestReq.BloodRequestID)
	if err != nil {
		logrus.Error("Could not find bloodrequest error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(bloodrequest)

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	BloodRequest storebloodrequest.BloodRequest
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &bloodrequestReader{
		bloodrequestes: params.BloodRequest,
	}
}
