package organization

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/organization/dto"
	"gitlab.com/Aubichol/blood-bank-backend/store/organization"
	storeorganization "gitlab.com/Aubichol/blood-bank-backend/store/organization"
	"go.uber.org/dig"
)

//Reader provides an interface for reading organizationes
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//organizationReader implements Reader interface
type organizationReader struct {
	organizationes organization.Organizations
}

func (read *organizationReader) askStore(organizationID string) (
	organization *model.Organization,
	err error,
) {
	organization, err = read.organizationes.FindByID(organizationID)
	return
}

func (read *organizationReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *organizationReader) prepareResponse(
	organization *model.Organization,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(organization)
	return
}

func (read *organizationReader) Read(organizationReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	organization, err := read.askStore(organizationReq.OrganizationID)
	if err != nil {
		logrus.Error("Could not find organization error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(organization)

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Organization storeorganization.Organizations
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &organizationReader{
		organizationes: params.Organization,
	}
}
