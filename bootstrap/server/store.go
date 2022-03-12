package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	donor "gitlab.com/Aubichol/blood-bank-backend/store/donor/mongo"
	notice "gitlab.com/Aubichol/blood-bank-backend/store/notice/mongo"
	patient "gitlab.com/Aubichol/blood-bank-backend/store/patient/mongo"
	staticcontent "gitlab.com/Aubichol/blood-bank-backend/store/staticcontent/mongo"

	//	picture "gitlab.com/Aubichol/blood-bank-backend/store/picture/mongo"
	//	status "gitlab.com/Aubichol/blood-bank-backend/store/status/mongo"
	//	token "gitlab.com/Aubichol/blood-bank-backend/store/token/mongo"
	user "gitlab.com/Aubichol/blood-bank-backend/store/user/mongo"
)

//Store provides constructors for mongo db implementations
func Store(c container.Container) {
	c.Register(patient.Store)
	c.Register(donor.Store)
	c.Register(staticcontent.Store)
	c.Register(notice.Store)
	//	c.Register(picture.Store)
	//	c.Register(status.Store)
	//	c.Register(token.Store)
	//	c.Register(comment.Store)
	c.Register(user.Store)
}
