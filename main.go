package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Vertex struct {
	x int
	y int
}

// User return array
type User []struct {
	UserID  int
	ID      int
	Caption string `json:"title" xml:"title"`
	Body    string
}

type customer struct {
	ID string
}

func (cus customer) String() {
	fmt.Sprintf("thanakorn" + cus.ID)
}

func main() {
	// v := Vertex{1, 2}
	// fmt.Println("TEST", v)
	// fmt.Println(v.x)
	// fmt.Println(v.y)
	var cus customer
	cus.ID = "1111"
	cus.String()
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(fetch("https://jsonplaceholder.typicode.com/posts"))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func fetch(url string) User {
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
