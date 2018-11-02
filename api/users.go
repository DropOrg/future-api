package api

import (
	"encoding/json"
	"hash"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var mg *mongo.Client

// User represents a user
type User struct {
	Username    string  `json:"id"`
	Name        string  `json:"name,omitempty"`
	Projects    []int64 `json:"projects,omitempty"`
	DateCreated int64   `json:"-"`
	password    hash.Hash
}

// CreateUser takes in a User object and returns an empty response body
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var b []byte
	r.Body.Read(b)
	err := json.Unmarshal(b, &user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
}

// ReadUser takes in { username: string } and returns a User
func ReadUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser takes in a User object and returns an empty response body.
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser takes in { username: string } and returns an empty response body.
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// UsersRoute creates and returns an instance of the router for all /users routes
func UsersRoute(route *mux.Route) *mux.Router {
	r := route.Subrouter()
	return r
}

// UserRoute creates and returns an instance of the router for all /user routes
func UserRoute(route *mux.Route) *mux.Router {
	r := route.Subrouter()
	r.Path("/{user}").
		Methods("GET").HandlerFunc(ReadUser).
		Methods("PUT").HandlerFunc(CreateUser).
		Methods("POST").HandlerFunc(UpdateUser).
		Methods("DELETE").HandlerFunc(DeleteUser)
	return r
}
