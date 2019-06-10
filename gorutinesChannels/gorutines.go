//gorutyny nie stosuje się w bibliotekach/packegach, współbieżność concurrency zostawia się do użytku dla użytkownika, niech programista korzystający z packegu sam zdecyduje, kiedy zechce użyć współbieżności
//trzeba wiedziec kiedy dana gorutyna sie skonczy, inaczej moze dzialac  tle caly czas i zajmowac pamiec, ale scrashowac po czasie program
//TODO: go run -race main.go   - wyświetli nam info o data race, jeśli np dwie gorutyny próboją dorwać się do tej same zmiennej w jednym czasie.
package main

import (
	"runtime"
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}

///////////////

func main() {
	go sayHello()
}

//nic się nie wydrukuje, func main jest wykonywana od razy i gorutyna nie ma czsu by sie wykonac

///////////////////

func hello2() {
	go sayHello()
	fmt.Println("Bye")
}

//gorutyna ma szanse sie wykonać, ale na 90% nie zdaży tego zrobić

/////////////////

func hello3() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

//tutaj gorutyna zdazy sie wykonać, bo program sobie na coś czeka

///////////////

func changingVariableAfterGorutine() {
	text := "Hello"
	go func() { //anonimowa funckja nie przyjmuje żadnych argumentów
		fmt.Println(text)
	}() //puste miejsce na argumenty
	text = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

//wydrukuje się "Goodbye", bo gorutyna czeka i wykonuje się dopiero po zmianie wartości text na goodbye. funkcja nie wykonuje sie od razu
//nie jest to gwarantowane, jest szansa, że wydrukuje się Hello

/////////////////////

func changingVariableAfterGorutine2() {
	text := "Hello"
	go func(argument string) {
		fmt.Println(argument)
	}(text) //zmienna text zostaje argumentem tej anonimowej funkcji
	text = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

//w 100% przypadków wydrukuje się Hello, bo argument do gorutyny został przekazany wczesniej, nawet jesli gorutyna wykonała się po zmianie zmiennej na inną wartość

//WAIT GROUPS
//służą do grupowania gorutyny w grupy, które mają sie wykonywać w danej kolejności

var wg = sync.WaitGroup{} //inicjalizacja wait group
var counter = 0

func sayHello() {
	fmt.Printf("hello #%v\n", counter)
	wg.Done() //funkcja gdy sie wykona daje o tym znać do wg
}
func increment() {
	counter++
	wg.Done() //funkcja gdy sie wykona daje o tym znać do wg
}
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(2)      //zadeklarowanie, ile gorutyn przyjmnie dana wait grupa
		go sayHello()  //gorutyna się wykonuje samodzielnie w danej wait grupie, dopóki nie dojdzie do wg.Done()
		go increment() //jw, nigdy nie bedzie tak, że najpierw liczba się podwyższy, a dopiero potem wydrukuje w danej petli
		wg.Wait()      //wait czeka na wszystkie gorutyny w danej grupie az sie wykonają
	}
}

//RWMutex
//służy do zamykania i otwierania wątków (??), dzieki niemu wykonuje się tylko jedna rzecz na raz
var counter = 0
var mutex = sync.RWMutex{}

func sayHello() {
	fmt.Printf("hello #%v\n", counter)
	mutex.RUnlock() //odblokowywuje odczyt
}
func increment() {
	counter++
	mutex.Unlock() //odblokowywuje zmiane danych
}
func main() {
	for i := 0; i < 100; i++ {
		mutex.RLock() //blokuje odczyt
		go sayHello()
		mutex.Lock() //blokuje zmiane danych
		go increment()
	}
}

// runtime.GOMAXPROCS()
//pozwala sprawdzic i/lub ustawić liczbę wątków, ktorą wykorzystuje dany program
runtime.GOMAXPROCS(-1) //zwraca liczbę działająych wątków
fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

runtime.GOMAXPROCS(50) //ustawi liczbę wątków na 50
//generalnie jeden wątek na rdzeń cpu to minimum, ale wiekszość aplikacji będzie szybciej działać na większej liczbie wątków
//trzeba to bezpośrednio przetestować i ustawić z juz prawie gotowym programem, by działał szybciej
//za dużo wątków sprawi że program bedzie działał wolniej