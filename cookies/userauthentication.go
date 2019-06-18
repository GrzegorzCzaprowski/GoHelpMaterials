package main

import (
	"encoding/json"
	"fmt"
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

	cookie := &http.Cookie{ //zainicjowanie
		Name:  "userlogcookie",
		Value: "0",
	}

	for _, u := range baza {
		if u == user {
			//val := fmt.Sprintf("true" + user.Name)
			cookie.Value = user.Name
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
	var cookie *http.Cookie
	cookie, err := req.Cookie("userlogcookie")
	if err != nil {
		fmt.Println(err)
		fmt.Println("cookie", cookie)
		return
	}
	name := cookie.Value
	cookie.Value = "false"
	cookie.MaxAge = -1
	http.SetCookie(w, cookie) //ustawienie ciastka dla odpowiedzi
	log.Printf("%v user logged out", name)
}

func change(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var cookie *http.Cookie
	cookie, err := req.Cookie("userlogcookie")
	if err != nil {
		fmt.Println(err)
		fmt.Println("cookie", cookie)
		return
	}
	userName := cookie.Value

	for i := range baza {
		if baza[i].Name == userName {

			baza[i].Logged = 111111

			log.Printf("%v user value changed %v", baza[i].Name, baza[i].Logged)
			return
		}
	}
	log.Printf("%v user cant be find", userName)

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
