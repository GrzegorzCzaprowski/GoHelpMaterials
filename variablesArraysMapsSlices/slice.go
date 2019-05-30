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