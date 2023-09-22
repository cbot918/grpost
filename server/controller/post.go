package controller

import (
	"context"
	"encoding/json"
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

type Comment struct {
	Id       string `json:"_id"`
	Text     string `json:"text"`
	PostedBy struct {
		Id   string `json:"_id"`
		Name string `json:"name"`
	}
}

type postResponse struct {
	Id        string    `json:"_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Comments  []Comment `json:"comments"`
	CreatedAt string    `json:"createdAt"`
	Likes     []string  `json:"likes"`
	Photo     string    `json:"photo"`
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
		fmt.Println(err)
		return c.Status(422).SendString(string("query posts failed"))
	}
	fmt.Println("posts : ", posts)

	comments, err := p.query.GetComments(context.Background())
	if err != nil {
		fmt.Println("comments not found")
		fmt.Println(err)
		return c.Status(422).SendString(string("query comments failed"))
	}

	likes, err := p.query.GetLikes(context.Background())
	if err != nil {
		fmt.Println("likes not found")
		fmt.Println(err)
		return c.Status(422).SendString(string("query likes failed"))
	}

	result := p.postWrapper(posts)

	result = p.commentWrapper(comments, result)

	result = p.likeWrapper(likes, result)

	return c.JSON(result)
}
func (p *Post) postWrapper(posts []db.Post) (resp []postResponse) {

	for _, item := range posts {

		fixedPost := postResponse{
			Id:        string(item.ID.String()),
			Title:     item.Title,
			Body:      item.Body,
			Comments:  []Comment{},
			CreatedAt: item.CreatedAt.String(),
			Likes:     []string{},
			Photo:     item.Photo,
			PostedBy: struct {
				Id   string `json:"_id"`
				Name string `json:"name"`
			}{
				Id:   item.PostedBy.String(),
				Name: item.UserName,
			},
			UpdatedAt: item.UpdatedAt.Time.String(),
		}

		resp = append(resp, fixedPost)
	}
	return
}

func (p *Post) commentWrapper(comments []db.GetCommentsRow, posts []postResponse) []postResponse {

	for _, c := range comments {
		for i := range posts {

			if posts[i].Id == c.TargetPost.String() {
				comment := Comment{
					Id:   c.ID.String(),
					Text: c.Texts,
					PostedBy: struct {
						Id   string `json:"_id"`
						Name string `json:"name"`
					}{
						Id:   c.PostedBy.String(),
						Name: c.UserName,
					},
				}

				posts[i].Comments = append(posts[i].Comments, comment)
			}

		}
	}
	return posts
}

func lj(v any) {
	res, err := json.MarshalIndent(v, "", "  ")
	if err != nil {

	}
	fmt.Println(string(res))
}

func (p *Post) likeWrapper(likes []db.PostLike, posts []postResponse) (updatedPosts []postResponse) {

	hashedPosts := make(map[string]postResponse)

	for _, post := range posts {
		hashedPosts[post.Id] = post
	}

	for _, like := range likes {
		if post, exists := hashedPosts[like.TargetPost.String()]; exists {
			fmt.Println(like.TargetPost.String())

			post.Likes = append(post.Likes, like.TargetPost.String())

			hashedPosts[like.TargetPost.String()] = post // Update the struct in the map

		}
	}

	for _, post := range hashedPosts {
		// if post.Id == "60ccb056-8e68-9018-44d0-929300000000" {
		// 	lj(post)
		// 	fmt.Println("gotch")
		// 	return nil
		// }
		updatedPosts = append(updatedPosts, post)
	}

	fmt.Printf("%#+v", updatedPosts)

	return updatedPosts
}
