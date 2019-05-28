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

//deklaracja tablicy
var a [10]int
var a [4][5]int
//przypisanie:
a[0] = 7
a[3][4] = 5


//deklaracja i przypisanie
var primes = [5]int{2,3,5,7,11}
primes := [5]int{2,3,5,7,11}
w dwuwymiarowej:
var dwuwymiarowa = [3][2]int{{1,2}, {3,4}, {5,6}}

//pozwolenie go na wywnioskowanie wielkosci tablicy
primes := [...]int{2,3,5,7,11}

//sprawdzenie wielkości tablicy:
len(primes)
//w dwuwymiarowej:
var dwuwymiarowa = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
var kolumna = dwuwymiarowa[0]
println(len(kolumna))
len(dwuwymiarowa[0]) //nie zadziała, musi to być osobna tablica

//uzycie range (nie potrzeba wtedy wielkosci sprawdzać):
primes := [5]int{2, 3, 5, 7, 11}
for index, value := range primes {
	fmt.Printf("index w tabeli: %d, liczba = %d\n", index, value)
}
/* wydrukuje:
index w tabeli: 0, liczba = 2
index w tabeli: 1, liczba = 3
index w tabeli: 2, liczba = 5
index w tabeli: 3, liczba = 7
index w tabeli: 4, liczba = 11
*/

//iterowanie po tablicy dwuwymiarowej:
const rzad = 3
const kolumna = 2
var dwuwymiarowa = [rzad][kolumna]int{{1, 2}, {3, 4}, {5, 6}}
for i := 0; i < rzad; i++ {
	for j := 0; j < kolumna; j++ {
		print(dwuwymiarowa[i][j])
	}
	print("\n")
}
/* wydrukuje:
12
34
56
*/

//iterowanie za pomocą range
//pierwsza zwracana wartosc to index (tutaj bez indeksu _ ), druga to kopia wartości elementu w tym indeksie
var dwuwymiarowa = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
for _, rzad := range dwuwymiarowa {
	for _, kolumna := range rzad {
		print(kolumna)
	}
	print("\n")
}
/* wydrukuje:
12
34
56
*/

//iterowanie dwuwymiarowej za pomocą range:
var dwuwymiarowa = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
for index1, rzad := range dwuwymiarowa {
	for index2, kolumna := range rzad {
		fmt.Printf("[%d][%d] : %d ", index1, index2, kolumna)
	}
	println()
}
/* wydrukuje:
[0][0] : 1 [0][1] : 2 
[1][0] : 3 [1][1] : 4 
[2][0] : 5 [2][1] : 6 
*/

//SLICE'Y 
//slicy same w sobie nie przechowuja zadnych danych, sa tylko odwołaniem do tablic
//pusty slice ma wartośc nil

//deklaracja slice'a z tablicy:
primes := [...]int{2, 3, 5, 7, 11} //tablica
var s []int		
s = primes[:1]	//slice tej tablicy
fmt.Println(s)
//wydrukuje:		[2]

s = primes[3:5]
//wydrukuje:	[7 11]

s = primes[2:]
//wydrukuje:		[5 7 11]

//slice literal:
var slicePrimes = []int{2, 3, 5, 7, 11}
slicePrimes := []int{2, 3, 5, 7, 11} //jak tablica, ale bez określonego rozmiaru

//tworzenie dynamicznego slice'a za pomocą make
sliceA := make([]int, 5) // wydrukuje: [0 0 0 0 0]

sliceB := make([]int, 0, 5)	//wydrukuje: []
fmt.Println(b)			
//trzeci argument make oznacza cap() capacity (pojemność) slicea, drugi aktualną dłougość len()
//gdy pojemność slice zostanie przekroczona, zostaje podwojona przez dwa 

//append dodawanie do slicea kolejnych wartości
var s []int		//nil slice, można appendowac do pustych slicow
s = append(s, 0)
s = append(s, 1)
s = append(s, 2)
// wydrukuje: [0 1 2]

//MAPY przypisują klucze do wartości
var m map[int]string  // mapa o wartości nil, nie da się przypisać klucza, nic z nią nie można zrobić poza i tak dalszym użyciem make

m := make(map[int]string)
m[2] = "major"		//wydrukuje się:	map[2:major]

//map literal:
var m = map[int]string{
	14: "Janusz",
	1234: "Marysia",
	3: "babajaga",
}

m[13] = "Dexter"        //wpisanie nowej pozycji do mapy
m[3] = "czarnoksieznik" //uzyto istniejacego klucza, nadpisanie obecnej pozycji w mapie
var wyciag = m[1234]    //wyciagniecie z mapy danego elementu i przypisanie go do zmiennej
delete(m, 1234)         //skasowanie elementu w mapie 			wydrukuje sie:

fmt.Println(m)			//wydrukuje:		map[3:czarnoksieznik 13:Dexter 14:Janusz]
fmt.Println(wyciag)		//					Marysia

value, test := m[3]			//sprawdzenie, czy do danego klucza przypisana i jest wartosc i wyciagniecie jej do zmiennej
fmt.Println(value, test)	//wydrukuje:	czarnoksieznik true
value, test = m[1234]		
fmt.Println(value, test)	//wydrukuje:	false	



//map literal ze structem
type Wspolrzedne struct {
rownoleznik	, poludnik float64
}
var m = map[string]Wspolrzedne{
	"Brama portowa": {40.68433, -74.39967},
	"Łobez":    {37.42202, -122.08408},
}
