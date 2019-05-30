// TODO: jeżeli usunie się jakiś objekt z mapy, to nie zostanie on tak naprawdę usunięty, jego wartości zostaną zastąpione
// nilami i zerami. PAMIEĆ NIE ZOSTANIE ZWOLNIONA. przy dużych mapach ryzyko braku pamięci (out of memory OOM)
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
