package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	Id      int    `json:"id"`
	Caption string `json:"caption"`
}

func main() {
	r := chi.NewRouter()

	fmt.Print(amishiPosts)
	r.Get("/posts", getPosts)

	http.ListenAndServe(":8888", r)
}

var amishiPosts []Post = []Post{
	{
		Id:      1,
		Caption: "aisha",
	},
	{
		Id:      2,
		Caption: "anesh",
	},
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Print("bye")

	Postbytes, err := json.Marshal(amishiPosts)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(Postbytes)
	w.Write(Postbytes)
}
