package staticcontent

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/pkg/tag"
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
	staticcontent staticcontent.StaticContent
}

func (read *staticcontentReader) askStore(staticcontentID string) (
	staticcontent *model.Status,
	err error,
) {
	staticcontent, err = read.staticcontentes.FindByID(staticcontentID)
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
	staticcontent *model.Status,
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

func (read *staticcontentReader) checkFriendShip(giverID, userID string) (
	currentRequest *model.FriendRequest,
	err error,
) {
	uniqueTag := tag.Unique(giverID, userID)
	currentRequest, err = read.friends.FindByUniqueTag(uniqueTag)
	return
}

func (read *staticcontentReader) Read(staticcontentReq *dto.ReadReq) (*dto.ReadResp, error) {
	//TO-DO: some validation on the input data is required
	staticcontent, err := read.askStore(staticcontentReq.StatusID)
	if err != nil {
		logrus.Error("Could not find staticcontent error : ", err)
		return nil, read.giveError()
	}

	var resp dto.ReadResp
	resp = read.prepareResponse(staticcontent)
	giverID := staticcontent.UserID
	//If the same person who has given the staticcontent asks for
	//the staticcontent, we should give them.
	if read.isSameUser(giverID, staticcontentReq.UserID) {
		return &resp, nil
	}

	currentRequest, err := read.checkFriendShip(
		giverID,
		staticcontentReq.UserID,
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
	Status  storestaticcontent.Status
	Friends friendrequest.FriendRequests
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &staticcontentReader{
		staticcontentes: params.Status,
		friends:         params.Friends,
	}
}
