package controller

import (
	"context"
	"fmt"

	db "github.com/cbot918/grpost/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	query *db.Queries
}

func NewPost(q *db.Queries) *Post {
	return &Post{
		query: q,
	}
}

type postRequest struct {
	Id       int32
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type postReqponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

func (p *Post) AllPost(c *fiber.Ctx) error {

	posts, err := p.query.GetPosts(context.Background())
	if err != nil {
		fmt.Println("posts not found")
		return c.Status(422).SendString(string("query posts failed"))
	}
	fmt.Println("posts : ", posts)

	return c.JSON(posts)
}
