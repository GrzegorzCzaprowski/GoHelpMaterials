//FUNKCJE
//funkcje to tez wartości i moga być przekazywane jak każda inna wartość (np jako wartość innej funkcji
//funckja może być zamknięta (closure) do jakiejś wartości
//nazwy metod z duzej litery są globalne, z małej sa prywatne
var Sum int				//zmienna globalna sum
func add1() int {		//funkcja add1 korzysta wartości która jest poza nią
	Sum++
	return Sum			//zwraca inta Sum
}
func main() {
	fmt.Println(add1())	//wydrukuje sie: 1
}


//VARIADIC FUNCTION
//funckja ze zmienną liczbą argumentów
func add (numbers ...int) int {		//trzykropek oznacza ze moge podać dowolną ilość argumentów danego typu
	sum := 0
	for _, value := range numbers{
		sum += value
	}
	return sum
}


//FUNCKAJ ANONIMOWA
var argument = 3
func (number int) int {	//funkcja bez nazwy
	return 3 * number
}(argument)					// argumenty funkcji podaje się w tym miejscu
//wydrukuje: 9


//FUNKCJA REKURENCYJNA
//funkcja która wywołuje sama siebie
func fact(n int) int {		//funkcja przyjmie jako n 3, potem 2 a na koncu 1, wywoła sie 3 razy po kolei
	if n == 1 {				//func(3) return 3 + func(2)
		return n			//func(2) return 2 + func(1)
	}						//func(1) return 1 
	return n + fact(n-1)	//jedna funkcja konczac się wykonywać wywołuje następną, i ta kolejną aż n będzie równe 1
}							//3+2+1
func main() {
	fmt.Println(fact(3))		//wydrukuje się 6
}


//METODY to funkcje przypisane do danych typów lokalnych, czyli:
//Wszystkie metody do danego typu powinny byc albo wartościami Type albo pointerami *Type
type Dupa struct{				//struct'ow
	waga int 
	kraglosc string
	kopnieta bool
}
type RozmiarButa int64		//nowy type ale o pojedynczej wartosci
int64 						//NIE DA SIĘ przypisac metody do typów prostych
func (d Dupa) kopnij(){
	d.kopnieta = true		//zmieni wartość kopnieta podanej d Dupy na true
}


//POINTERY, czyli zmienne wskaźnikowe pozwalających na bezpośredni dostęp do pamięci,
// pamięć jest reprezentowana jako jednowymiarowa tablica bajtów – wszystkie zmienne (statyczne i dynamiczne) są umieszczane w tej „tablicy”. 

//można deklarowac metody za pomocą pointer receiver'ow *T. pracują one wtedy bezpośrenio na danym typie, a nie jego kopii
//jak Value Receiver
func (d Dupa) schudnijTroche(){			//nie ma pointera, operujemy na kopii
	d.waga = d.waga - 5
}
func main() {
	d := Dupa{20, "wypukla", false}		//stworzenie dupy
	d.schudnijTroche()					//metoda operuje na kopii d, a nie zmienia orginalnej zmiennej
	fmt.Println(d)						//wydrukuje: 20
}

//jako Pointer Receiver
func (d *Dupa) schudnijTroche(){		//jest pointer, operujemy na orginale
	d.waga = d.waga - 5					
}
func main() {
	d := Dupa{20, "wypukla", false}		//stworzenie dupy
	d.schudnijTroche()					//metoda operuje na orginalnym d (wskazuje gdzie się znajduje w pamięci), 
	fmt.Println(d)						//nastapi zmiana w danych orginalnego sructa
}	//wydrukuje: 15

//pointery w funkcjach:
func schudnijTroche(d Dupa){			//nie ma pointera, operujemy na kopii structa
	d.waga = d.waga - 5
}
func main() {
	d := Dupa{20, "wypukla", false}    
	schudnijTroche(d)					//funkcja operuje na kopii d, a nie zmienia orginalnej zmiennej
	fmt.Println(d)						//zmiany nie nastąpią
}	//wydrukuje: 20

//z REFERENCJĄ, operatorem referencji, operatorem adresu - zwraca fizyczny adres komórki pamięci przechowującej &zmienną
//wartość, która zawiera informacje o położeniu innej wartości w pamięci 
func schudnijTroche(d *Dupa){			//nie ma pointera, operujemy na kopii structa
	d.waga = d.waga - 5
}
func main() {
	d := Dupa{20, "wypukla", false}    
	schudnijTroche(&d)					//funkcja operuje na orginale, wskazanym przez & operator adresu
	fmt.Println(d)						//zmiany nastąpią
}	//wydrukuje: 15