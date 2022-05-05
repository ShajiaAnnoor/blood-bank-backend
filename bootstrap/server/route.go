package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/api/bloodrequest"
	"gitlab.com/Aubichol/blood-bank-backend/api/donor"
	"gitlab.com/Aubichol/blood-bank-backend/api/notice"
	"gitlab.com/Aubichol/blood-bank-backend/api/organization"
	"gitlab.com/Aubichol/blood-bank-backend/api/patient"
	"gitlab.com/Aubichol/blood-bank-backend/api/staticcontent"
	"gitlab.com/Aubichol/blood-bank-backend/api/user"

	"gitlab.com/Aubichol/blood-bank-backend/api/wsroute"
	"gitlab.com/Aubichol/blood-bank-backend/container"
)

//Route registers all the route providers to container
func Route(c container.Container) {
	c.RegisterGroup(bloodrequest.CreateRoute, "route")
	c.RegisterGroup(bloodrequest.ReadRoute, "route")
	c.RegisterGroup(bloodrequest.UpdateRoute, "route")
	c.RegisterGroup(bloodrequest.DeleteRoute, "route")

	c.RegisterGroup(donor.CreateRoute, "route")
	c.RegisterGroup(donor.ReadRoute, "route")
	c.RegisterGroup(donor.UpdateRoute, "route")
	c.RegisterGroup(donor.DeleteRoute, "route")

	c.RegisterGroup(notice.CreateRoute, "route")
	c.RegisterGroup(notice.ReadRoute, "route")
	c.RegisterGroup(notice.UpdateRoute, "route")
	c.RegisterGroup(notice.DeleteRoute, "route")

	c.RegisterGroup(organization.CreateRoute, "route")
	c.RegisterGroup(organization.ReadRoute, "route")
	c.RegisterGroup(organization.UpdateRoute, "route")
	c.RegisterGroup(organization.DeleteRoute, "route")

	c.RegisterGroup(patient.CreateRoute, "route")
	c.RegisterGroup(patient.ReadRoute, "route")
	c.RegisterGroup(patient.UpdateRoute, "route")
	c.RegisterGroup(patient.DeleteRoute, "route")

	c.RegisterGroup(staticcontent.CreateRoute, "route")
	c.RegisterGroup(staticcontent.ReadRoute, "route")
	c.RegisterGroup(staticcontent.UpdateRoute, "route")
	c.RegisterGroup(staticcontent.DeleteRoute, "route")

	c.RegisterGroup(user.RegistrationRoute, "route")
	c.RegisterGroup(user.LoginRoute, "route")
	c.RegisterGroup(user.SearchRoute, "route")

	c.RegisterGroup(wsroute.WSRoute, "ws_route")
}
