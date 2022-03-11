package notice

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/notice/dto"
	storenotice "gitlab.com/Aubichol/blood-bank-backend/store/notice"
	"go.uber.org/dig"
)

//Reader provides an interface for reading noticees
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//noticeReader implements Reader interface
type noticeReader struct {
	notices storenotice.Notices
}

func (read *noticeReader) askStore(noticeID string) (
	notice *model.Notice,
	err error,
) {
	notice, err = read.notices.FindByID(noticeID)
	return
}

func (read *noticeReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *noticeReader) prepareResponse(
	notice *model.Notice,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(notice)
	return
}

func (read *noticeReader) isSameUser(giverID, userID string) (
	isSame bool,
) {
	isSame = giverID == userID
	return
}

func (read *noticeReader) Read(noticeReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	notice, err := read.askStore(noticeReq.NoticeID)
	if err != nil {
		logrus.Error("Could not find notice error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(notice)
	giverID := notice.UserID
	//If the same person who has given the notice asks for
	//the notice, we should give them.
	if read.isSameUser(giverID, noticeReq.UserID) {
		return &resp, nil
	}

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Notice storenotice.Notices
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &noticeReader{
		notices: params.Notice,
	}
}
