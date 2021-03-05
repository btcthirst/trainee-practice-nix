package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

//Articles is an exported struct
type Articles struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//Testcomments is an exported struct and named as db.table
type Testcomments struct {
	gorm.Model
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

const (
	host     = "127.0.0.1"
	port     = ":3306"
	database = "golang"
	user     = "webmax"
	password = "qwerty1234#WebM4X"
)

func dbConn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection faild")
	}
	fmt.Println("connection ok")
	return db
}

func ins(com []Testcomments) {
	db := dbConn()

	db.Create(&com)

}

func getDataJSON(urla string, c string) {
	resp, err := http.Get(urla)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	go unmarshJSON(body, c)

}

func unmarshJSON(b []byte, check string) {
	var posts []Articles
	var comments []Testcomments
	switch check {
	case "posts":
		json.Unmarshal(b, &posts)
		for _, v := range posts {

			urla := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", v.ID)

			go getDataJSON(urla, "comments")

		}
	case "comments":
		json.Unmarshal(b, &comments)

		go ins(comments)

	default:
		fmt.Println("I don't know how to do this yet. contact the developers")
	}

}

func main() {

	uIdent := 7
	check := "posts"
	urla := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", uIdent)

	go getDataJSON(urla, check)

	fmt.Scan(&check)
	fmt.Println(check)

}
