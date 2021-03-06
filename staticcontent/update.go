package staticcontent

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent/dto"
	storestaticcontent "gitlab.com/Aubichol/blood-bank-backend/store/staticcontent"
	validator "gopkg.in/go-playground/validator.v9"
)

//Updater provides an interface for updating staticcontentes
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates user staticcontent
type update struct {
	storeStaticContent storestaticcontent.StaticContents
	validate           *validator.Validate
}

func (u *update) toModel(userstaticcontent *dto.Update) (sc *model.StaticContent) {
	sc = &model.StaticContent{}
	sc.UpdatedAt = time.Now().UTC()
	sc.Text = userstaticcontent.StaticContent
	sc.UserID = userstaticcontent.UserID
	sc.ID = userstaticcontent.StaticContentID
	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelStaticContent *model.StaticContent,
) {
	modelStaticContent = u.toModel(update)
	return
}

func (u *update) askStore(modelStaticContent *model.StaticContent) (
	id string,
	err error,
) {
	id, err = u.storeStaticContent.Save(modelStaticContent)
	return
}

func (u *update) giveResponse(
	modelStaticContent *model.StaticContent,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelStaticContent.UserID,
	}).Debug("User updated staticcontent successfully")

	return &dto.UpdateResponse{
		Message:    "Static Content updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelStaticContent.UpdatedAt.String(),
	}
}

func (u *update) giveError() (err error) {
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "Invalid data",
		},
	}
	err = fmt.Errorf(
		"%s %w",
		err.Error(),
		&errResp,
	)
	return
}

//Update implements Update interface
func (u *update) Update(update *dto.Update) (
	*dto.UpdateResponse, error,
) {
	if err := u.validateData(update); err != nil {
		return nil, err
	}

	modelStaticContent := u.convertData(update)
	fmt.Println("id is", modelStaticContent.ID)
	id, err := u.askStore(modelStaticContent)
	if err == nil {
		return u.giveResponse(modelStaticContent, id), nil
	}

	logrus.Error("Could not update staticcontent ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of NewCreate
func NewUpdate(store storestaticcontent.StaticContents, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
