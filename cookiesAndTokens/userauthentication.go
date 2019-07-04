package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Content  string `json:"content,omitempty"`
}

var baza []User

// var key = []byte("Nothing to see here, goy.")

// func hashCookie(cookievalue string) string {
// 	h := hmac.New(sha256.New, key)
// 	io.WriteString(h, cookievalue)
// 	return fmt.Sprintf("%x", h.Sum(nil))
// }

// func checkCookie(hashedCookieValue string) bool {
// 	return hmac.Equal([]byte(hashedCookieValue), key)
// }

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
		if u.Name == user.Name && u.Password == user.Password {

			cookie.Value = user.Name
			cookie.MaxAge = 30
			http.SetCookie(w, cookie) //ustawienie ciastka dla odpowiedzi
			log.Printf("%v user Logged\n", user.Name)
			return
		}
	}
	cookie.Value = "false"
	cookie.MaxAge = -1
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
	log.Printf("%v user Password out", name)
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

			baza[i].Content = "111111"

			log.Printf("%v user value changed %v", baza[i].Name, baza[i].Password)
			return
		}
	}
	log.Printf("%v user cant be find", userName)

}

func main() {
	router := httprouter.New()

	user1 := User{Name: "user1", Password: "3", Content: "12345"}
	baza = append(baza, user1)

	router.POST("/post", post)
	router.GET("/getall", getall)
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/change", change)

	log.Fatal(http.ListenAndServe(":8000", router))
}
