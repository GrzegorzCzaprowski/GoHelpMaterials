package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	makeGetRequest()
	makePostRequest()
	makePostRequestWithHeader()
}

func makeGetRequest() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func makePostRequest() {

	jsonData := map[string]string{"name": "greg", "lastname": "bush"} //stworzenie mapy danych
	jsonValue, _ := json.Marshal(jsonData)                            //zamiana danych na forma jsona

	//zbuforowanie jasona do bytów
	response, err := http.Post("http://httpbin.org/post", "aplication/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}

func makePostRequestWithHeader() {

	jsonData := map[string]string{"name": "greg", "lastname": "bush"} //stworzenie mapy danych
	jsonValue, _ := json.Marshal(jsonData)                            //zamiana danych na forma jsona

	//zbuforowanie jasona do bytów
	request, _ := http.NewRequest("POST", "http://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}
