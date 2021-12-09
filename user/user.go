package user

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"strconv"
	"user-cli/repo"
)

func AddUserToDataBase(user *User) {
	db := repo.GetRepo()
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}

	var b *bolt.Bucket
	b = tx.Bucket([]byte("users"))
	if b == nil {
		b, err = tx.CreateBucket([]byte("users"))
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	userId, err := b.NextSequence()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	user.UserId = &userId

	buf, err := json.Marshal(user)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	} else {
		userIdStr := strconv.FormatUint(*user.UserId, 10)
		log.Println(userIdStr)
		err := b.Put([]byte(userIdStr), buf)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	return
}

func GetUsers(userId *uint64) []*User {
	db := repo.GetRepo()
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}

	userList := make([]*User, 0)
	b := tx.Bucket([]byte("users"))
	if b == nil {
		return userList
	}
	if userId != nil {
		userID := fmt.Sprintf(strconv.FormatUint(*userId, 10))
		userBytes := b.Get([]byte(userID))
		if len(string(userBytes)) == 0 {
			return userList
		}
		var user User
		err = json.Unmarshal(userBytes, &user)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		userList = append(userList, &user)
		return userList
	}

	cursor := b.Cursor()
	for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
		var user User
		err := json.Unmarshal(v, &user)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		userList = append(userList, &user)
	}
	tx.Commit()
	return userList
}
