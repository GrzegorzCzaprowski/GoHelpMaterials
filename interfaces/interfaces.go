//INTERFACE'Y
//interfacy to też rodzaj wartości jak int czy jakis struct
//jedna zmienna może miec wiecej niz jeden typ czyli zmienna burek jest typu dog i animal
//interfejsy nie opisują danych (od tego structy), tylo opisują zachowania
// w structach znajdują sie dane, zmienne, w interfejsach są metody

//lepiej uzywac wielu malych interfejsów, niż jednego dużego

type Dupa interface {
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
type Massager interface {
	Massage()
}
type Dupsko interface {
	Kicker
	Massager
}

//jezeli w interface'ie jest metoda przyjmująca pointer receiver *Type, to implementacja tego
//interfejsu także musi odbywać się przez pointer. Jeżeli metoda przyjmuje value, to interfejs implementuje sie normalnie
// pusty interface to najwiekszy zbiór rzeczy w GO
//"interface{} says nothing"

//typ moze implementować interfejs poprzez uzycie jego funckji. Nie trzeba bezpośrenio implementować, nie ma słowka kluczowego "implements"
type Massager interface {
	Massage()
}
type Dypa struct {
	wymasowana bool
}

func (d Dupa) Massage() {
	//cialo funckji
}

//Empty interfaces are used by code that handles values of unknown type. 
//For example, fmt.Print takes any number of arguments of type interface{}. 

//type assertions zapewnia, że jakas wartosc var o typie interface bedzie przetrzymywać konkretny typ. 
//Mozna sprawdzic czy interface trzyma konretny Type, assertion wroci wtedy daną wartośc i booleana czy zapewnienie się udało
var inter interface{} = "hello"

	s := inter.(string)
	fmt.Println(s)		//Wydrukuje: hello

	s, ok := inter.(string)
	fmt.Println(s, ok)		//wydrukuje: hello true

	f, ok := inter.(float64)	//wydrukuje: 0 false
	fmt.Println(f, ok)

	f = inter.(float64) //wyjebie panica
	fmt.Println(f)

//TYPE SWITCH
//konstrukcja, ktora robi okreslone rzeczy zaleznie od typu danego interfejsu
func do(i interface{}) {			//funckja do przyjmuje jako typ danych interfejs
	switch v := i.(type) {			//sprawdza jakiego typu jest interfejs, deklaracja taka sama jak w type assertion, tylko zamiast konkretnego Typu używamy słówka kluczego "type"
	case int:				
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	do(21)			//wydrukuje: Twice 21 is 42
	do("hello")		//wydrukuje: "hello" is 5 bytes long
	do(true)		//wydrukuje: I don't know about type bool!
}