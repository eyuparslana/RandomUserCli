package repo

import (
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

var RandomUser Repo

type Repo struct {
	db *bolt.DB
}

func init() {
	db, err := bolt.Open("/tmp/random-user.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	RandomUser.db = db
}

func GetRepo() *bolt.DB {
	return RandomUser.db
}
