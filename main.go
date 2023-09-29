package main

import (
	"errors"
	"log"

	"fsedano.net/test/db"
)

func main() {
	checkErrors(db.DbOpen("nonexistingdb"))
	checkErrors(db.DbOpen("notpermissionsdb"))
	checkErrors(db.DbOpen("otherbrokenthings"))

}

func checkErrors(err error) {
	if err != nil {
		var dbErr *db.DbError
		var nonexist *db.NonExistingDbError
		var noperm *db.NonPermissionsDbError
		switch {
		case errors.As(err, &dbErr):
			log.Printf("It was a db err!: %s", err.Error())
		case errors.As(err, &nonexist):
			log.Printf("It didnt exist: %s", err.Error())
		case errors.As(err, &noperm):
			log.Printf("We didnt have permissions: %s", err.Error())
		default:
			log.Printf("A generic error happened: %s", err.Error())
		}
	}

}
