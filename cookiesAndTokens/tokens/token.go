package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"isadmin,omitempty"`
}

type Claims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isadmin"`
	jwt.StandardClaims
}

var jwtKey = []byte("nothing to see here goy")

var DB []User

func main() {
	router := httprouter.New()

	DB = append(DB, User{Name: "user", Password: "user", IsAdmin: false})
	DB = append(DB, User{Name: "admin", Password: "admin", IsAdmin: true})

	router.GET("/public", public)
	router.POST("/login", login)
	router.POST("/private", private)
	router.POST("/admin", admin)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func public(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	log.Println("info for everyone")
}

func login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("error with decoding user to json: ", err)
		w.WriteHeader(500)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Name,
		IsAdmin:  user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, u := range DB {
		if user.Name == u.Name && user.Password == u.Password && user.IsAdmin == u.IsAdmin {
			log.Println("loggin succesfull")
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
			return
		} else if user.Name == u.Name && user.Password != u.Password {
			log.Println("bad password!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	log.Println("bad password or username")
	w.WriteHeader(http.StatusUnauthorized)

}

func private(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("no token")
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	log.Println("you are logged and have authoriztion to this site")
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}

func admin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if claims.IsAdmin == true {
		// Finally, return the welcome message to the user, along with their
		// username given in the token
		log.Println("welcome admin")
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("you are not an admin")
	}
}
