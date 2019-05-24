package main

// poradnik z tąd: https://youtu.be/t96hBT53S4U

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"` //wyluczy jesli bedzie pusty
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person //globalna zmienna, która może byc użyta w kązdej z pozostałych funkcji

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)       //tutaj są trzymane parametry z endpointa
	for _, item := range people { //iteruje po całym slice people
		if item.ID == params["id"] { //sprawdza parametr id z danego endpointa(url-a), jak znajdzie to wychodzi z funkcji
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{}) //zwraca pustą osobe jeśli nie znajdzie osoby o podanym Id w forze wyżej
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people) //wyswietli wszystkich ludzi na liscie
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) //tutaj są trzymane parametry z endpointa
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person) //bierze json date z Postbody i marshaluje ją do zmiennej person
	person.ID = params["id"]                      //ustawia id zmiennej person na to co zostało przekazane w parametrze
	/*bierze ip z params (czyli numer ip przekazany w urlu,
	np http://localhost:8000/people/3) i wstawia w pole Id structa person Person*/
	people = append(people, person)   //dołoncza nowo stworzoną osobę do slicea people
	json.NewEncoder(w).Encode(people) //wyswietli wszystkich ludzi na liscie po wykonaniu sie kodu który jest wyżej
}

func CreatePersonFromUrlEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)       //tutaj są trzymane parametry z endpointa
	var person Person             //zainicjalizowanie zmiennej typu Person by było do czego przypisać parametry
	person.Address = new(Address) // TODO: WAŻNE: zainicjowanie zmiennej typu Address, która nie inicjuje się automatycznie przy incjalizacji zmiennej typu Person
	person.ID = params["id"]      //przypisywanie konkretnych parametrow z urla do konkretnych pol konkretnego person
	person.Firstname = params["firstname"]
	person.Lastname = params["lastname"]
	person.Address.City = params["city"]
	person.Address.State = params["state"]
	people = append(people, person)   //dołoncza nowo stworzoną osobę do slicea people
	json.NewEncoder(w).Encode(people) //wyswietli wszystkich ludzi na liscie po wykonaniu sie kodu który jest wyżej
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)           //tutaj są trzymane parametry z endpointa
	for index, item := range people { //leci po całym slice people
		if item.ID == params["id"] { //aż znajdzie podane w parametrze id
			people = append(people[:index], people[index+1:]...) //usunięcie osoby polega na wzieciu wszystkiego co jest przed
			break                                                //szukanym id, i wszystkiego co jest po i ponownym złaczeniu tych dwóch części
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	/*ruter, definicja z wiki jako urządzenia: służy do łączenia różnych sieci komputerowych
	(różnych w sensie informatycznym, czyli np. o różnych klasach, maskach itd.), pełni więc rolę węzła komunikacyjnego.
	Na podstawie informacji zawartych w pakietach TCP/IP jest w stanie przekazać pakiety z dołączonej
	do siebie sieci źródłowej do docelowej, rozróżniając ją spośród wielu dołączonych do siebie sieci.*/

	//mock data:
	people = append(people, Person{ID: "1", Firstname: "grzegorz", Lastname: "Czaprowski", Address: &Address{City: "Szczecin", State: "ZachodnioPomorSKIEM"}})
	people = append(people, Person{ID: "2", Firstname: "Jan", Lastname: "Czaprowski"})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	//moje metody
	router.HandleFunc("/people/{id}/{firstname}/{lastname}/{city}/{state}", CreatePersonFromUrlEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
