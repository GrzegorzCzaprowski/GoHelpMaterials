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