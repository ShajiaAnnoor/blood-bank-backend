package patient

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/notice/dto"
	storepatient "gitlab.com/Aubichol/blood-bank-backend/store/patient"
	"go.uber.org/dig"
)

//Reader provides an interface for reading patients
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//patientReader implements Reader interface
type patientReader struct {
	patients patientnotice.Notice
}

func (read *patientReader) askStore(patientID string) (
	patient *model.Patient,
	err error,
) {
	patient, err = read.patients.FindByID(patientID)
	return
}

func (read *patientReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *patientReader) prepareResponse(
	patient *model.Patient,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(patient)
	return
}

func (read *patientReader) Read(patientReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	patient, err := read.askStore(patientReq.PatientID)
	if err != nil {
		logrus.Error("Could not find patient error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(patient)
	giverID := patient.UserID
	//If the same person who has given the patient asks for
	//the patient, we should give them.
	if read.isSameUser(giverID, patientReq.UserID) {
		return &resp, nil
	}

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Patient storepatient.Patient
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &patientReader{
		patients: params.Patient,
	}
}
