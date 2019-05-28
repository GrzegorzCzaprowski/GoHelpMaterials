package main

import (
	"fmt"
)

//uzywane w momentach krytycznych, ktore jak sie zjebia to aplikacja sie wysypie, errory nie sa tak wazne, po errorach moze dalej dzialać
//w sytuacjach nieodwracalnych (np podzielenie przez zero, bład przy wczytaniu pliku który jest kluczowy dla programu)

func basicPanic() {
	fmt.Println("start")
	panic("something bad happen")
	fmt.Println("end")
	/*wydrukuje się:
	start
	panic: something bad happen
	+pierdoly, reszta kodu się nie wykona

	w go error to cos normalnego, panic wystepuje gdy kod już całkiem nie może kontynuować
	cały kod pod paniciem nie wykona się
	*/
}

/////////////////

//schemat kolejności:
func panicAndDeferDependence() {
	+fmt.Println("start")                  //normalne isntrukcje przed paniciem sie wykonaja
	defer fmt.Println("this was deferred") //zdeferowane instrukcje przed paniciem beda czekac az otaczajaca je funkcja sie wykona
	panic("something bad happen")          // panic ubija funkcje, ale zanim to zrobi, wszystkie zdeferowane wczsesniej instrukcje się wykonują
	fmt.Println("end")                     // instrukcje po panicu w danej funckcji juz się nie wykonaja
	defer fmt.Println("end deferred")      // zdeferowane też nie

	/*wydrukuje się:
	start
	this was deferred
	panic: something bad happen
	+pierdoly, reszta kodu się nie wykona

	panic wyskakuje dopiero po wykonaniu się instrukcji zdeferowanych*/

}

///////////////////

func main() {
	panicAndDeferDependence()

}
