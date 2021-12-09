package user

import (
	"time"
)

const (
	BASEURL = "https://randomuser.me/api/"
)

var OperatorMap = map[string]func(limit, age int) bool{
	"e":   func(limit, age int) bool { return limit == age },
	"gt":  func(limit, age int) bool { return limit < age },
	"gte": func(limit, age int) bool { return limit <= age },
	"lt":  func(limit, age int) bool { return limit > age },
	"lte": func(limit, age int) bool { return limit >= age },
}

type Params struct {
	Gender      string
	Nationality string
	Count       string
}

type Filter struct {
	Age         int
	AgeOperator string
	Nationality string
	Gender      string
	UserId      uint64
}

type Response struct {
	Results []*User `json:"results"`
}
type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}
type Dob struct {
	Date time.Time `json:"date"`
	Age  int       `json:"age"`
}
type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}
type User struct {
	UserId  *uint64 `json:"user_id"`
	Gender  string  `json:"gender"`
	Name    Name    `json:"name"`
	Email   string  `json:"email"`
	Dob     Dob     `json:"dob"`
	Picture Picture `json:"picture"`
	Nat     string  `json:"nat"`
}
