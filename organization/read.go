package organization

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/organization/dto"
	"gitlab.com/Aubichol/blood-bank-backend/pkg/tag"
	storeorganization "gitlab.com/Aubichol/blood-bank-backend/store/organization"
	"go.uber.org/dig"
)

//Reader provides an interface for reading organizationes
type Reader interface {
	Read(*dto.ReadReq) (*dto.ReadResp, error)
}

//organizationReader implements Reader interface
type organizationReader struct {
	organizationes organization.Notice
	friends        friendrequest.FriendRequests
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

func (read *organizationReader) isSameUser(giverID, userID string) (
	isSame bool,
) {
	isSame = giverID == userID
	return
}

func (read *organizationReader) checkFriendShip(giverID, userID string) (
	currentRequest *model.FriendRequest,
	err error,
) {
	uniqueTag := tag.Unique(giverID, userID)
	currentRequest, err = read.friends.FindByUniqueTag(uniqueTag)
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
	giverID := organization.UserID
	//If the same person who has given the organization asks for
	//the organization, we should give them.
	if read.isSameUser(giverID, organizationReq.UserID) {
		return &resp, nil
	}

	currentRequest, err := read.checkFriendShip(
		giverID,
		organizationReq.UserID,
	)
	if err != nil {
		logrus.Error("Could not find friendship error : ", err)
		return nil, read.giveError()
	}

	if currentRequest.Organization != "accepted" {
		return nil, err
	}

	return &resp, nil
}

//NewReaderParams lists params for the NewReader
type NewReaderParams struct {
	dig.In
	Organization storeorganization.Organization
	Friends      friendrequest.FriendRequests
}

//NewReader provides Reader
func NewReader(params NewReaderParams) Reader {
	return &organizationReader{
		organizationes: params.Organization,
		friends:        params.Friends,
	}
}
