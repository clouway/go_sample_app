package main

import (
	"log"
	"net/http"

	"github.com/mgenov/myproject/api"
	"github.com/mgenov/myproject/persistence"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	sessionStore := persistence.NewSessionStore(session)

	http.Handle("/", api.Adapt(api.IndexHandler(),
		api.CheckAuth(sessionStore),
	))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
