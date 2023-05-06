package config

import (
	r "github.com/dancannon/gorethink"
	"log"
)

func getSession() *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "whoosh",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	return session
}
