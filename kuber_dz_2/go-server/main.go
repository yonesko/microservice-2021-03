/*
 * User Service
 *
 * This is simple client API
 *
 * API version: 1.0.0
 * Contact: schetinnikov@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/yonesko/microservice-2021-03/kuber_dz_2/go-server/db"
	sw "github.com/yonesko/microservice-2021-03/kuber_dz_2/go-server/go"
)

func main() {
	repo := initRepo()
	_, err := repo.FindUser(5)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", sw.NewRouter()))
}

func initRepo() db.Repo {
	connStr := fmt.Sprintf("user=%s dbname=%s", mustEnv("PG_USER"), mustEnv("PG_DB"))
	dat, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db.PgRepo{Db: dat}
}

func mustEnv(s string) string {
	e := os.Getenv(s)
	if len(s) == 0 {
		panic("no " + s + " in ENV")
	}
	return e
}
