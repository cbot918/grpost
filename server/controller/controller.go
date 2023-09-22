package controller

import db "github.com/cbot918/grpost/db/sqlc"

type Controller struct {
	Query *db.Queries
	Auth  *Auth
	Post  *Post
}

func NewController(query *db.Queries) *Controller {

	ctlr := new(Controller)
	ctlr.Auth = NewAuth(query)
	ctlr.Post = NewPost(query)
	ctlr.Query = query

	return ctlr
}
