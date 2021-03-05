package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// page for create articles
func createPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/createPage.html", "./templates/header.html", "./templates/footer.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "create", nil)
}

//page for update/delete articles
func updatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/updatePage.html", "./templates/header.html", "./templates/footer.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "update", nil)
}

//page for create comments
func commentsPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/commentsPage.html", "./templates/header.html", "./templates/footer.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "comments", nil)
}

//main page
func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/homePage.html", "./templates/header.html", "./templates/footer.html", "./templates/sidebar.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "home", nil)
}

func initServ() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/create/", createPage)
	http.HandleFunc("/update/", updatePage)
	http.HandleFunc("/comments/", commentsPage)

	fmt.Println(http.ListenAndServe(":7575", nil))

}

func main() {
	initServ()
	fmt.Println("test")

}
