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
		return getErrorRes(c, err, "invalid email or password")
		// return c.Status(422).SendString(err.Error())
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

	dbuser, err := a.query.CreateUser(context.Background(), db.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		fmt.Println("db create user failed")
		// return getErrorRes(c, "no user")
		return c.Status(http.StatusBadRequest).SendString(string("error when db create user"))
	}

	res := signupResponse{
		Email:   dbuser.Email,
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

func getErrorRes(c *fiber.Ctx, err error, message string) error {

	data := struct {
		Err    error  `json:"error"`
		Mesage string `json:"message"`
	}{}
	data.Err = err
	fmt.Println(data)
	data.Mesage = message
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal json error")
	}

	return c.Status(422).SendString(string(d))

}
