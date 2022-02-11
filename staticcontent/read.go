package staticcontent

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/notice/dto"
	"gitlab.com/Aubichol/blood-bank-backend/pkg/tag"
	storestatus "gitlab.com/Aubichol/blood-bank-backend/store/notice"
	"go.uber.org/dig"
)

//Reader provides an interface for reading statuses
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//statusReader implements Reader interface
type statusReader struct {
	statuses status.Notice
	friends  friendrequest.FriendRequests
}

func (read *statusReader) askStore(statusID string) (
	status *model.Status,
	err error,
) {
	status, err = read.statuses.FindByID(statusID)
	return
}

func (read *statusReader) giveError() (err error) {
	err = &errors.Unknown{
		errors.Base{
			"Invalid request", false,
		},
	}
	return
}

func (read *statusReader) prepareResponse(
	status *model.Status,
) (
	resp dto.ReadResp,
) {
	resp.FromModel(status)
	return
}

func (read *statusReader) isSameUser(giverID, userID string) (
	isSame bool,
) {
	isSame = giverID == userID
	return
}

func (read *statusReader) checkFriendShip(giverID, userID string) (
	currentRequest *model.FriendRequest,
	err error,
) {
	uniqueTag := tag.Unique(giverID, userID)
	currentRequest, err = read.friends.FindByUniqueTag(uniqueTag)
	return
}

func (read *statusReader) Read(statusReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	status, err := read.askStore(statusReq.StatusID)
	if err != nil {
		logrus.Error("Could not find status error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(status)
	giverID := status.UserID
	//If the same person who has given the status asks for
	//the status, we should give them.
	if read.isSameUser(giverID, statusReq.UserID) {
		return &resp, nil
	}

	currentRequest, err := read.checkFriendShip(
		giverID,
		statusReq.UserID,
	)
	if err != nil {
		logrus.Error("Could not find friendship error : ", err)
		return nil, read.giveError()
	}

	if currentRequest.Status != "accepted" {
		return nil, err
	}

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Status  storestatus.Status
	Friends friendrequest.FriendRequests
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &statusReader{
		statuses: params.Status,
		friends:  params.Friends,
	}
}
