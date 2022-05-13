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

//Deleter provides an interface for updating staticcontentes
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes staticcontent
type delete struct {
	storeStaticcontent storestaticcontent.StaticContents
	validate           *validator.Validate
}

func (d *delete) toModel(userstaticcontent *dto.Delete) (staticcontent *model.StaticContent) {
	staticcontent = &model.StaticContent{}
	staticcontent.UpdatedAt = time.Now().UTC()
	staticcontent.UserID = userstaticcontent.UserID
	staticcontent.ID = userstaticcontent.StaticContentID
	staticcontent.IsDeleted = true
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modelStaticcontent *model.StaticContent,
) {
	modelStaticcontent = d.toModel(delete)
	return
}

func (d *delete) askStore(modelStaticcontent *model.StaticContent) (
	id string,
	err error,
) {
	id, err = d.storeStaticcontent.Save(modelStaticcontent)
	return
}

func (d *delete) giveResponse(
	modelStaticcontent *model.StaticContent,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelStaticcontent.UserID,
	}).Debug("User deleted staticcontent successfully")

	return &dto.DeleteResponse{
		Message: "Staticcontent deleted",
		OK:      true,
		ID:      id,
		//		DeleteTime: modelStaticcontent.DeletedAt.String(),
	}
}

func (d *delete) giveError() (err error) {
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

//Delete implements Delete interface
func (d *delete) Delete(delete *dto.Delete) (
	*dto.DeleteResponse, error,
) {
	if err := d.validateData(delete); err != nil {
		return nil, err
	}

	modelStaticcontent := d.convertData(delete)
	id, err := d.askStore(modelStaticcontent)
	if err == nil {
		return d.giveResponse(modelStaticcontent, id), nil
	}

	logrus.Error("Could not delete staticcontent ", err)
	err = d.giveError()
	return nil, err
}

//NewDelete returns new instance of NewDelete
func NewDelete(store storestaticcontent.StaticContents, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
