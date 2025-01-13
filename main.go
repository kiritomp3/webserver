package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gavruk/go-blog-example/models"
)

var posts map[string]*models.Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)

}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	post := models.NewPost(id, title, content)
	posts[post.Id] = post
}
func main() {
	fmt.Println("listening on port:8080")

	posts = make(map[string]*models.Post, 0)

	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/SavePost", savePostHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.ListenAndServe(":8080", nil)
}
