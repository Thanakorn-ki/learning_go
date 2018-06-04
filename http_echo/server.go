package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// User return array
type User struct {
	UserID  int
	ID      int
	Caption string `json:"title" xml:"title"`
	Body    string
}

// Users is array
type Users []User

func main() {
	e := echo.New()
	e.GET("/posts", getUsers)
	e.GET("/posts/:id", getUser)
	e.GET("/insert/", connectMongoDB)
	e.Logger.Fatal(e.Start(":1323"))
}

func connectMongoDB(c echo.Context) error {
	client, err := mongo.NewClient("mongodb://salapao2136:root1234@ds147180.mlab.com:47180/mongo-golang")
	collection := client.Database("mongo-golang").Collection("users")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	res, err := collection.InsertOne(context.Background(), map[string]string{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	mid := fmt.Sprintf("%v", id)
	return c.JSON(http.StatusOK, mid)
}
func getUsers(c echo.Context) error {
	data := fetch("https://jsonplaceholder.typicode.com/posts")
	return c.JSON(200, data)
	// return c.XML(200, data)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	data := fetchUser("https://jsonplaceholder.typicode.com/posts/" + id)
	// return c.JSON(200, data)
	return c.XML(200, data)
}

func fetch(url string) Users {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "146070bf-9b3c-a8e1-a0a1-96ef04c2460b")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var u Users
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Fatal(err)
	}

	return u
}

func fetchUser(url string) User {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "146070bf-9b3c-a8e1-a0a1-96ef04c2460b")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var u User
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Fatal(err)
	}

	return u
}
