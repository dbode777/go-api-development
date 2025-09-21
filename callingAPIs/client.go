package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: Refactor to reference project's structs package
type User struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const BASE_URL = "http://localhost:8080"

func getUsersService(endpoint string) []byte {
	response, err := http.Get(BASE_URL + endpoint)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return responseData
}

func getUsers() {
	responseData := getUsersService("/users")
	var users []User
	json.Unmarshal(responseData, &users)
	fmt.Println(users) // This would be where you would either call a DB or another API service
}

func getUserById(id string) {
	responseData := getUsersService("/users/" + id)
	var user User
	json.Unmarshal(responseData, &user)
	fmt.Println(user) // This would be where you would either call a DB or another API service
}

func main() {
	getUsers()
	getUserById("1")
}
