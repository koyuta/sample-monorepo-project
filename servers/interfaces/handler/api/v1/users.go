package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"servers/domain"
	"servers/interfaces/handler"
	"servers/pkg/router"
	"servers/usecase"
)

type Users struct {
	router router.Router
	users  *usecase.Users
}

func NewUsers(r router.Router, u *usecase.Users) *Users {
	return &Users{router: r, users: u}
}

type UsersGetResponse struct {
	ID int64 `json:"id"`
}

func (u *UsersGetResponse) Marshal() (*domain.User, error) {
	user := &domain.User{ID: u.ID}
	return user, nil
}

func (u *Users) Get(rw http.ResponseWriter, r *http.Request) {
	var uid int64
	if id, err := strconv.ParseUint(u.router.URLParam(r, "id"), 10, 32); err == nil {
		uid = int64(id)
	}
	if uid == 0 {
		handler.WriteNotFoundToHeader(rw, handler.NewErrorResponse("Not found"))
		return
	}

	user, err := u.users.Get(context.Background(), uid)
	if err != nil {
		// TODO: Write a error log to logger.
		handler.WriteInternalServerErrorToHeader(rw, handler.NewErrorResponse("some error occurred"))
		return
	}
	if user == nil {
		handler.WriteNotFoundToHeader(rw, handler.NewErrorResponse("Not found"))
		return
	}
	handler.WriteOKToHeader(rw, user)
}

type UsersPostRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (u *UsersPostRequest) Unmarshal(d *domain.User) error {
	if u.Name == "" {
		return errors.New("name filed is required")
	}
	d.Name = u.Name
	d.Description = u.Description
	return nil
}

func (u *Users) Post(rw http.ResponseWriter, r *http.Request) {
	var req = &UsersPostRequest{}
	if err := handler.RetrieveBody(r.Body, req); err != nil {
		handler.WriteBadRequestToHeader(rw, handler.NewErrorResponse("invalid"))
		return
	}

	var user *domain.User
	if err := req.Unmarshal(user); err != nil {
		handler.WriteBadRequestToHeader(rw, handler.NewErrorResponse(err.Error()))
		return
	}

	id, err := u.users.Add(context.Background(), user)
	if err != nil {
		// TODO: Write a error log to logger.
		handler.WriteInternalServerErrorToHeader(rw, handler.NewErrorResponse("some error occurred"))
		return
	}

	handler.WriteCreatedToHeader(rw, path.Join("/users", fmt.Sprint(id)))
}
