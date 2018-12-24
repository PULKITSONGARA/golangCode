package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

const (
	mongoUrl = "localhost:27017"
)

// type Account struct {
// 	gorm.Model
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	Token    string `json:"token";sql:"-"`
// }

type Person struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
}

func login(w http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// user_id := req.FormValue("user_id")
	// password := req.FormValue("password")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))
	var t Person
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	user_id := t.User_id
	password := t.Password
	c := session.DB("DataX").C("user_details")
	res := Person{}
	err = c.Find(`{"user_id": user_id}`).One(&res)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user_id,
		"password": password,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	// fmt.Println(res.user_id, err)
	fmt.Print(password)
	fmt.Print(tokenString)
	w.Write([]byte(tokenString))
}
