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