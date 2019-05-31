//INTERFACE'Y
//interfacy to też rodzaj wartości jak int czy jakis struct
//jedna zmienna może miec wiecej niz jeden typ czyli zmienna burek jest typu dog i animal
//interfejsy nie opisują danych (od tego structy), tylo opisują zachowania
// w structach znajdują sie dane, zmienne, w interfejsach są metody

//lepiej uzywac wielu malych interfejsów, niż jednego dużego

type Dupa interface{
	Kopnij()
}
//w przypadku interfejsu z jedną metodą, interfejs powininien nazywac się jak ta metoda +er
type Kicker interface {
	Kick()
}

//interfejsy można dodawac do innych interfejsów (tak jak structy)
type Kicker interface {
	Kick()
}
type Massager interface{
	Massage()
}
type Dupsko interface{
	Kicker
	Massager
}

//jezeli w interface'ie jest metoda przyjmująca pointer receiver *Type, to implementacja tego 
//interfejsu także musi odbywać się przez pointer. Jeżeli metoda przyjmuje value, to interfejs implementuje sie normalnie
// pusty interface to najwiekszy zbiór rzeczy w GO
//"interface{} says nothing"