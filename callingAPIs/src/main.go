package main

import (
	errorhandler "example/callingAPIs/src/errorHandler"
	structs "example/callingAPIs/src/structs"

	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []structs.User{
	{Id: "1", Name: "John", Age: 30},
	{Id: "2", Name: "Jane", Age: 25},
	{Id: "3", Name: "Bob", Age: 40},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(id string) (*structs.User, error) {
	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errorhandler.NewError(errorhandler.ErrUserNotFound)
}

func userById(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Issue retrieving user data", "error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/users", getUsers)
	router.GET("/users/:id", userById)
	router.Run("localhost:8080")
}
