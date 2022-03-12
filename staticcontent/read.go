package staticcontent

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent/dto"
	storestaticcontent "gitlab.com/Aubichol/blood-bank-backend/store/staticcontent"
	"go.uber.org/dig"
)

//Reader provides an interface for reading staticcontentes
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//staticcontentReader implements Reader interface
type staticcontentReader struct {
	staticcontent storestaticcontent.StaticContents
}

func (read *staticcontentReader) askStore(staticcontentID string) (
	staticcontent *model.StaticContent,
	err error,
) {
	staticcontent, err = read.staticcontent.FindByID(staticcontentID)
	return
}

func (read *staticcontentReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *staticcontentReader) prepareResponse(
	staticcontent *model.StaticContent,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(staticcontent)
	return
}

func (read *staticcontentReader) isSameUser(giverID, userID string) (
	isSame bool,
) {
	isSame = giverID == userID
	return
}

func (read *staticcontentReader) Read(staticcontentReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	staticcontent, err := read.askStore(staticcontentReq.StaticContentID)
	if err != nil {
		logrus.Error("Could not find staticcontent error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(staticcontent)
	//If the same person who has given the staticcontent asks for
	//the staticcontent, we should give them.

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	StaticContent storestaticcontent.StaticContents
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &staticcontentReader{
		staticcontent: params.StaticContent,
	}
}
