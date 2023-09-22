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

/*
	{
		"id":"",
		"title":"",
		"body":"",
		"comments":[],
		"createdAt":"",
		"likes":[],
		"photo":"",
		"postedBy":{
			"_id":"",
			"name":""
		},
		"updatedAt":""
	}
*/

/*
​​ID: "64208a79-b4cb-dc00-1e06-4a1a00000000"
​​Title: "蟲蟲剋星"
Body: "想起有個BUG放了一年好像還是不會解 :D"
comments:
​​CreatedAt: "2023-03-26T18:10:01.65Z"
likes:
​​Photo: "http://res.cloudinary.com/yalecloud/image/upload/v1679854200/ktnlvnns4zat4lv8wakk.png"
​​PostedBy: "60ccae67-fcd2-c732-9ca3-96c000000000"

​​UpdatedAt: Object { Time: "2023-03-29T15:41:15.26Z", Valid: true }

*/

type postReqponse struct {
	Id        string   `json:"_id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Comments  []string `json:"comments"`
	CreatedAt string   `json:"createdAt"`
	Likes     []string `json:"likes"`
	Photo     string   `json:"photo"`
	PostedBy  struct {
		Id   string `json:"_id"`
		Name string `json:"name"`
	}
	UpdatedAt string `json:"updatedAt"`
}

func (p *Post) AllPost(c *fiber.Ctx) error {

	posts, err := p.query.GetPosts(context.Background())
	if err != nil {
		fmt.Println("posts not found")
		return c.Status(422).SendString(string("query posts failed"))
	}
	fmt.Println("posts : ", posts)

	fn := posts[49]
	fmt.Println(fn.ID)

	// mockpost := postReqponse{
	// 	Id:        string(fn.ID.String()),
	// 	Title:     fn.Title,
	// 	Body:      fn.Body,
	// 	Comments:  []string{},
	// 	CreatedAt: fn.CreatedAt.String(),
	// 	Likes:     []string{},
	// 	Photo:     fn.Photo,
	// 	PostedBy: struct {
	// 		Id   string `json:"_id"`
	// 		Name string `json:"name"`
	// 	}{
	// 		Id:   "user123",
	// 		Name: "yale mock",
	// 	},
	// 	UpdatedAt: fn.UpdatedAt.Time.String(),
	// }

	// mockposts := []postReqponse{}

	// mockposts = append(mockposts, mockpost)

	return c.JSON(p.postWrapper(posts))
}

func (p *Post) postWrapper(posts []db.Post) (resp []postReqponse) {

	for _, item := range posts {

		fixedPost := postReqponse{
			Id:        string(item.ID.String()),
			Title:     item.Title,
			Body:      item.Body,
			Comments:  []string{},
			CreatedAt: item.CreatedAt.String(),
			Likes:     []string{},
			Photo:     item.Photo,
			PostedBy: struct {
				Id   string `json:"_id"`
				Name string `json:"name"`
			}{
				Id:   "user123",
				Name: "yale mock",
			},
			UpdatedAt: item.UpdatedAt.Time.String(),
		}

		resp = append(resp, fixedPost)
	}
	return
}
