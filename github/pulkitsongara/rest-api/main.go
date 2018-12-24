// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// var people []Person

// type Person struct {
// 	ID        string   `json:"id,omitempty"`
// 	Firstname string   `json:"firstname,omitempty"`
// 	Lastname  string   `json:"lastname,omitempty"`
// 	Address   *Address `json:"address,omitempty"`
// }
// type Address struct {
// 	City  string `json:"city,omitempty"`
// 	State string `json:"state,omitempty"`
// }

// func GetPeople(w http.ResponseWriter, r *http.Request) {

// 	params := mux.Vars(r)
// 	for _, item := range people {
// 		if item.ID == params["id"] {
// 			fmt.Println("if done")
// 		}
// 	}
// 	json.NewEncoder(w).Encode(people)
// 	fmt.Println("hello API hit")
// }

// // our main function
// func main() {
// 	people := append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
// 	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
// 	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

// 	router := mux.NewRouter()

// 	router.HandleFunc("/people", GetPeople).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8000", router))
// }

// // return "doneeeee"

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

const (
	mongoUrl = "localhost:27017"
)

type Person struct {
	Name  string
	Phone string
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/insert", insert).Methods("POST")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
	// router.HandleFunc("/list", GetPeople).Methods("GET")
	// router.HandleFunc("/delete", GetPeople).Methods("DELETE")
	// router.HandleFunc("/update", GetPeople).Methods("PATCH")

	// userService := mongo.NewUserService(session.Copy(), dbName, userCollectionName)

}

func insert(w http.ResponseWriter, req *http.Request) {
	// vars := mux.Vars(r)
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	dbName := req.FormValue("dbName")
	CollectionName := req.FormValue("CollectionName")
	c := session.DB(dbName).C(CollectionName)
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"})
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("DONE"))

}

// func update() {

// }
