package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID        int    `json:"id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName "`
}

var persons = []person{
	{ID: 1, FirstName: "cristino", LastName: "ronaldo"},
	{ID: 2, FirstName: "sergio", LastName: "aquero"},
	{ID: 3, FirstName: "john", LastName: "cena"},
}

func getPerson(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, persons)

}

func addPerson(context *gin.Context) {
	var newPerson person

	if err := context.BindJSON(&newPerson); err != nil {
		return
	}

	persons = append(persons, newPerson)

	context.IndentedJSON(http.StatusCreated, newPerson)

}

func getPersons(context *gin.Context) {
	id := context.Param("id")
	person, err := getPersonById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, person)
}

func togglePersonStatus(context *gin.Context) {
	id := context.Param("id")
	person, err := getPersonById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
		return

	}

	context.IndentedJSON(http.StatusOK, person)

}

func getPersonById(id string) (*person, error) {
	for i, p := range persons {
		if p.ID == p.ID {
			return &persons[i], nil
		}

	}
	return nil, errors.New("person not found")

}

func healthCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, healthCheck)
	fmt.Printf("api is up and running")
}

func main() {
	router := gin.Default()
	router.GET("/persons", getPerson)
	router.GET("/healthcheck", healthCheck)
	router.GET("/persons/:id", getPersons)
	router.PATCH("/persons/:id", togglePersonStatus)
	router.POST("/persons", addPerson)
	router.Run("localhost:9090")
}
