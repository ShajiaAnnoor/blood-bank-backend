package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	bloodrequest "gitlab.com/Aubichol/blood-bank-backend/store/bloodrequest/mongo"
	donor "gitlab.com/Aubichol/blood-bank-backend/store/donor/mongo"
	notice "gitlab.com/Aubichol/blood-bank-backend/store/notice/mongo"
	organization "gitlab.com/Aubichol/blood-bank-backend/store/organization/mongo"
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
	c.Register(user.Store)
	c.Register(bloodrequest.Store)
	c.Register(organization.Store)
}
