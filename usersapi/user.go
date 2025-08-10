package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type users struct {
	UserName   string `json:"username"`
	UserID     int    `json:"userid"`
	YearJoined int    `json:"yearjoined"`
}

func main() {
	r := chi.NewRouter()
	x := 5
	fmt.Print(x)
	var y int = 5 
	fmt.Print(y)
	r.Get("/users", getUsers)
	r.Delete("/users", deleteUsers)
	r.Put("/users", updateUsers)
	r.Post("/users", createUsers)
	num := 42


    var ptr *int = &num


    fmt.Println(ptr)      
    fmt.Println(*ptr)  
    *ptr = 100
    fmt.Println(num)
	fmt.Println("Server running at http://localhost:8888")
	http.ListenAndServe(":8808", r)
	
}

var userslist []users = []users{
	{
		UserName:   "amishi",
		UserID:     1,
		YearJoined: 2000,
	},
	{
		UserName:   "aneesh",
		UserID:     2,
		YearJoined: 1980,
	},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	Postbytes, err := json.Marshal(userslist)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(Postbytes)
	w.Write(Postbytes)

}
type deleteUserRequestBody struct {
	Id int `json:"id"`

}
func deleteUsers(w http.ResponseWriter, r *http.Request) {
	var requestBody deleteUserRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
	for i, user := range userslist {
		if user.UserID == requestBody.Id {
			userslist = append(userslist[:i], userslist[i+1:]...)
		}
	}
	w.Write([]byte("updated successfully"))

}
type updateRequestBody struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
func updateUsers(w http.ResponseWriter, r *http.Request) {
	var requestBody updateRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err!= nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
	for i := range userslist {
		if userslist[i].UserID == requestBody.Id {
			userslist[i].UserName = requestBody.Name
		}
	}
	
	w.Write([]byte("updated successfully"))
}
type createRequestBody struct {
	Id int `json:"id"` 
	Name string `json:"name"`
	YearJoined int `json:"yearjoined"`
}
func createUsers(w http.ResponseWriter, r *http.Request) {
	var requestBody createRequestBody 
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}
	var newuser users = users{
		UserID: requestBody.Id,
		UserName: requestBody.Name,
		YearJoined: requestBody.YearJoined,
	}
	userslist = append(userslist, newuser)
	w.Write([]byte("added user"))
}