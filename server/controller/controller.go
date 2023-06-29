package controller

import db "github.com/cbot918/grpost/db/sqlc"

type Controller struct {
	Query *db.Queries
	Auth  *Auth
}

func NewController(query *db.Queries) *Controller {

	ctlr := new(Controller)
	ctlr.Auth = NewAuth(query)
	ctlr.Query = query

	return ctlr
}
