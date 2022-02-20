package user

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	storeuser "gitlab.com/Aubichol/blood-bank-backend/store/user"
	"gitlab.com/Aubichol/blood-bank-backend/user/dto"
)

//Searcher defines an interface for searching a user
type Searcher interface {
	Search(text string) ([]*dto.Search, error)
}

//SearcherFunc is a function type that implements Searcher interface
type SearcherFunc func(text string) ([]*dto.Search, error)

//Search implements Searcher interface
func (s SearcherFunc) Search(text string) ([]*dto.Search, error) {
	return s(text)
}

//NewSearcher provides Searcher
func NewSearcher(storeUsers storeuser.Users) Searcher {
	f := func(text string) ([]*dto.Search, error) {
		users, err := storeUsers.Search(text, int64(0), int64(20))
		if err != nil {
			logrus.WithField("text", text).Error("could not search user")
			return nil, fmt.Errorf("%s:%w", err.Error(), &errors.Unknown{
				Base: errors.Base{"invalid request", false},
			})
		}

		dtos := []*dto.Search{}
		for _, user := range users {
			searchDto := dto.Search{}
			searchDto.FromModel(user)
			dtos = append(dtos, &searchDto)
		}

		return dtos, nil
	}

	return SearcherFunc(f)
}
