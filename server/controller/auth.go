package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	db "github.com/cbot918/grpost/db/sqlc"
	"github.com/cbot918/liby/jwty"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	query *db.Queries
}

func NewAuth(q *db.Queries) *Auth {
	return &Auth{
		query: q,
	}
}

type signinRequest struct {
	Id       int32
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signinReqponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

func (a *Auth) Signin(c *fiber.Ctx) error {

	// get request
	user := signinRequest{}
	user.Id = 1
	if err := c.BodyParser(&user); err != nil {
		fmt.Println("err via bodyparser")
		return err
	}

	// db search user
	ctx := context.Background()
	target, err := a.query.GetUser(ctx, user.Email)
	if err != nil {
		fmt.Println("user not found")
		return c.Status(http.StatusBadRequest).SendString("user not found")
	}
	fmt.Println("user found: ", target)

	// jwt process
	token, err := jwty.New().FastJwt(int(user.Id), user.Email)
	if err != nil {
		fmt.Println("err call FastJwt")
		panic(err)
	}
	res := signinReqponse{}
	res.Token = token
	res.User = strings.Trim(regexp.MustCompile(".*@").FindString(user.Email), "@")

	// res response
	resp, err := json.Marshal(res)
	if err != nil {
		fmt.Println("err via json Marshal")
	}

	return c.SendString(string(resp))
}

type signupRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type signupResponse struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (a *Auth) Signup(c *fiber.Ctx) error {
	req := signupRequest{}
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("failed in bodyparser")
		return err
	}

	if a.userExist(req.Email) {
		return c.Status(http.StatusBadRequest).SendString("email is exist, please try another one")
	}

	a.query.CreateUser(context.Background(), db.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})

	res := signupResponse{
		Email:   req.Email,
		Message: "signup success",
	}
	resp, err := json.Marshal(res)
	if err != nil {
		fmt.Println("err via json Marshal")
	}
	return c.SendString(string(resp))
}

func (a *Auth) userExist(email string) bool {
	_, err := a.query.GetUser(context.Background(), email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println("faild in db get user")
		return false
	}
	return true
}
