package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name   string `json:"name,omitempty"`
	Logged int    `json:"logged,omitempty"`
}

var baza []User

func post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := User{}
	json.NewDecoder(req.Body).Decode(&user)

	baza = append(baza, user)
	log.Println("user added")
}

func getall(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(baza)
	log.Println("users getted")
}

func login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := User{}
	json.NewDecoder(req.Body).Decode(&user)

	cookie := &http.Cookie{ //zadeklarowanie ciastka
		Name: "userlogcookie",
	}

	for _, u := range baza {
		if u == user {
			cookie.Value = "true"
			http.SetCookie(w, cookie) //ustawienie ciastka dla odpowiedzi
			log.Printf("%v user logged\n", user.Name)
			return
		}
	}
	cookie.Value = "false"
	http.SetCookie(w, cookie) //ustawienie ciastka dla odpowiedzi
	log.Printf("%v cant login\n", user.Name)
}

func logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

func change(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

func main() {
	router := httprouter.New()

	user1 := User{Name: "user1", Logged: 3}
	baza = append(baza, user1)

	router.POST("/post", post)
	router.GET("/getall", getall)
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/change", change)

	log.Fatal(http.ListenAndServe(":8000", router))
}
