package main

import (
	"github.com/couchbase/gocb"
	"encoding/json"
	"fmt"
)

type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Social []SocialMedia `json:"socialmedia,omitempty"`
}

type SocialMedia struct {
	Title string `json:"title"`
	Link string `json:"link"`
}

func main() {
	cluster,_ := gocb.Connect("couchbase://localhost")
	bucket,_  := cluster.OpenBucket("default", "")

	var person Person

	bucket.Upsert("12345", Person{
		Firstname: "Daniel Sergiyovich",
		Lastname: "Sokyrynskyi",
		Social: []SocialMedia{
			{Title:"GitHub", Link: "https://github.com/dsokyrynskyi"},
			{Title:"Gmail", Link: "sokirinskiy@gmail.com"},
		},
	}, 0)

	/*bucket.Get("12345", &person)
	jsonBytes, _ := json.Marshal(person)
	fmt.Println(string(jsonBytes))*/

	query := gocb.NewN1qlQuery("SELECT default.* FROM default WHERE META().id = $1")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"12345"})
	rows.One(&person)
	jsonBytes, _ := json.Marshal(person)
	fmt.Println(string(jsonBytes))
}
