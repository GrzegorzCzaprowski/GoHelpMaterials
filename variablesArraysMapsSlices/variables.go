//LITERAŁ LITERAL: określona, niezmienna wartość w kodzie, której nie można zmienić. 
//czyli defacto każda cyfra w kodzie, nie można zmienić wartości liczby "2" na coś innego, dwójka zawsze będzie dwójką
x := 5 			// 5 jest literałem
y := "tekst" 	// "tekst" jest literałem
z := time()/2 	// literałem jest tu tylko 2, ponieważ jest ustaloną wartością
z := time()/x	 // $x nie jest literałem, (posiada wartość, ale niekoniecznie musi ona być zawsze taka sama)

//deklaracja zmiennej:
var x int
var x, y string

//deklaracja i definicja zmiennej:
var number = 2  //samo przypisze typ
number := 2
var number int = 2
var number = int(2)
number := int(2)

//deklaracja i definicja wielu zmiennych:
var a, b, c int = 1, 2, 3
var slowo, liczba = "nibemben", 3
slowo, liczba := "nibemben", 3
var (
	firstName, lastName string
	age                 int
)

//jasna konwersja typów:
var a int64 = 4
var b int = int(a)