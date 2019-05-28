package main

import "fmt"

func main() {

	var m = map[int]string{
		14:   "Janusz",
		1234: "Marysia",
		3:    "babajaga",
	}

	m[13] = "Dexter"        //wpisanie nowej pozycji do mapy
	m[3] = "czarnoksieznik" //uzyto istniejacego klucza, nadpisanie obecnej pozycji w mapie
	var wyciag = m[1234]    //wyciagniecie z mapy danego elementu i przypisanie go do zmiennej
	delete(m, 1234)         //skasowanie elementu w mapie
	value, test := m[3]
	fmt.Println(value, test)
	value, test = m[1234]
	fmt.Println(value, test)
	fmt.Println(m)
	fmt.Println(wyciag)

}
