package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Articles is an exported struct
type Articles struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//Comments is an exported struct
type Comments struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

const (
	host     = "127.0.0.1"
	database = "golang"
	user     = "webmax"
	password = "qwerty1234#WebM4X"
)

func conn(urla string, c string) {
	conn, err := http.Get(urla)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(conn.Body)
	if err != nil {
		panic(err)
	}

	go unmarshJ(body, c)

}

func unmarshJ(b []byte, check string) {
	var posts []Articles
	var comm []Comments
	switch check {
	case "posts":
		json.Unmarshal(b, &posts)
		for _, v := range posts {

			urla := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", v.ID)

			go conn(urla, "comments")

		}
	case "comments":
		json.Unmarshal(b, &comm)
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

		go conDb(connectionString, comm)

	default:
		fmt.Println("I don't know how to do this yet. contact the developers")
	}

}

func conDb(name string, comm []Comments) {

	for _, vc := range comm {
		db, err := sql.Open("mysql", name)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insData := fmt.Sprintf("INSERT INTO testcomments (postid, id, name, email, body) VALUES('%d', '%d', '%s', '%s', '%s')", vc.PostID, vc.ID, vc.Name, vc.Email, vc.Body)

		_, err = db.Exec(insData)
		if err != nil {
			panic(err)
		}

	}

}

func main() {

	uIdent := 7
	check := "posts"
	urla := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", uIdent)

	go conn(urla, check)

	fmt.Scan(&check)
	fmt.Println(check)
}
