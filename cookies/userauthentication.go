package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

type User struct {
	Name   string
	Logged int
}

func post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := User{}
	json.NewDecoder(req.Body).Decode(&user)

	baza = append(Baza, user)
}

func login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

func logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

func change(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

var baza []User

func main() {
	router := httprouter.New()

	user1 := User{Name: "user1", Logged: 3}
	baza = append(baza, user1)

	router.POST("/postuser", post)
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/change", change)

	log.Fatal(http.ListenAndServe(":8000", router))
}
