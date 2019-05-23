package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	a := "start"
	defer fmt.Println(a)
	a = "end"
	//wydrukuje się start, zmiana wartości po zdeferowaniu instrukcji nie zmienia wartości w tej instrukcji

	////////////////

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	/*wydrukuje się:

	start
	end
	middle

	zdeferowana instrukcja wyskakuje na końcu */

	///////////////

	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	/*wydrukuje się:

	end
	middle
	start

	zdeferowany kod wykonuje się  w odwrotnej kolejności, LIFO (last-in, first-out)  */

	///////////////

	resource, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	//defer sprawia, że zasób zostanie zamknięty na koniec funkcji w której się znajduje, pomimmo, że został zapisany tuż za nią
	defer resource.Body.Close()

	text, err := ioutil.ReadAll(resource.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", text)

	//defer nie nadaje się do for-ów i funkcji które otwierają naraz wiele zasobów, bo nie zostaną one zamknięte aż funkcja nie zostanie wykonana

	////////////////

}
