package openapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/types"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) PostUser(w http.ResponseWriter, r *http.Request) {
	created := CreateDummyUser()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
func (s *Server) GetUsersUserId(w http.ResponseWriter, r *http.Request, userId int) {
	created := CreateDummyUser()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
func (s *Server) PutUsersUserId(w http.ResponseWriter, r *http.Request, userId int) {
	created := CreateDummyUser()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
func (s *Server) DeleteUsersUserId(w http.ResponseWriter, r *http.Request, userId int) {
	created := CreateDummyUser()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func CreateDummyUser() *User {
	now := time.Now()
	id := 123456

	return &User{
		CreateDate:  &types.Date{now},
		DateOfBirth: types.Date{now},
		Email:       "test@example.com",
		FirstName:   "First",
		Id:          &id,
		LastName:    "Last",
	}
}
