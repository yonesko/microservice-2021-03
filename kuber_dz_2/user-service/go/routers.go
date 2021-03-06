/*
 * User Service
 *
 * This is simple client API
 *
 * API version: 1.0.0
 * Contact: schetinnikov@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(repo Repo) *mux.Router {
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/api/v1/").
		Name("Index").
		Handler(Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, "Hello World!")
		}), "Index"))
	router.
		Methods("POST").
		Path("/api/v1/user").
		Name("CreateUser").
		Handler(createUser(repo))
	router.
		Methods("DELETE").
		Path("/api/v1/user/{userId}").
		Name("DeleteUser").
		Handler(deleteUser(repo))
	router.
		Methods("GET").
		Path("/api/v1/user/{userId}").
		Name("FindUserById").
		Handler(findUser(repo))
	router.
		Methods("PUT").
		Path("/api/v1/user/{userId}").
		Name("UpdateUser").
		Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		}))

	return router
}

func createUser(repo Repo) http.Handler {
	return Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(400)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		id, err := repo.SaveUser(user)
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprint(id)))
	}), "createUser")
}

func findUser(repo Repo) http.Handler {
	return Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/user/"))
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		user, err := repo.FindUser(int64(id))
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		bytes, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(bytes)
	}), "findUser")
}
func deleteUser(repo Repo) http.Handler {
	return Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/user/"))
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		err = repo.DeleteUser(int64(id))
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}), "deleteUser")
}
