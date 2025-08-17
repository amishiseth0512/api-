package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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
	req, err := http.NewRequest("GET", os.Getenv("BASE_URL")+"/rest/v1/users", nil)
	if err != nil {
		fmt.Print(err)
		return
	}
	req.Header.Add("apikey", os.Getenv("API_KEY"))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_KEY"))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}
	fmt.Print(body)
	w.Write(body)

}

type deleteUserRequestBody struct {
	Id int `json:"id"`
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	var reqBody deleteUserRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest("DELETE", os.Getenv("BASE_URL")+fmt.Sprintf("/rest/v1/users?id=eq.%d", reqBody.Id), nil)
	if err != nil {
		fmt.Print(err)
		return
	}
	req.Header.Add("apikey", os.Getenv("API_KEY"))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_KEY"))
	client := http.Client{}
	resp, err := client.Do(req)
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println("body")
	if err != nil {
		fmt.Print(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(err)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}
	w.Write([]byte("deleted successfully"))

}

type updateRequestBody struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	var reqBody updateRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	payload := map[string]string{
		"name":       reqBody.Name,
		"created_at": time.Now().Format(time.RFC3339),
	}
	b, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	bodyReader := bytes.NewReader(b)
	req, err := http.NewRequest("PATCH", os.Getenv("BASE_URL")+fmt.Sprintf("/rest/v1/users?id=eq.%d", reqBody.Id), bodyReader)
	if err != nil {
		fmt.Print(err)
		return
	}
	req.Header.Add("apikey", os.Getenv("API_KEY"))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_KEY"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println("body")
	if err != nil {
		fmt.Print(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(err)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}
	w.Write([]byte("updated successfully"))
}

type createRequestBody struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	YearJoined int    `json:"yearjoined"`
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	var reqBody createRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}
	payload := map[string]string{
		"name":       reqBody.Name,
		"created_at": time.Now().Format(time.RFC3339),
	}
	b, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	bodyReader := bytes.NewReader(b)
	req, err := http.NewRequest("POST", os.Getenv("BASE_URL")+"/rest/v1/users", bodyReader)
	fmt.Println("err")
	fmt.Println(err)
	if err != nil {
		fmt.Print(err)
		return
	}

	req.Header.Add("apikey", os.Getenv("API_KEY"))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("API_KEY"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println("body")
	if err != nil {
		fmt.Print(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(err)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	w.Write([]byte("added user"))
}
