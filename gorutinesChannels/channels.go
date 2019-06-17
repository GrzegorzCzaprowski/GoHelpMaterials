//inicjalizacja kanałów, inaczej się nie da chyba
channel := make(chan int)	//kanał dwukierunkowy, może odbierać i wysyłać dane
ReceiveOnlyChannel := make(<-chan int)		//dane są pobierane z kanału
SendOnlyChannel := make(chan<- int)			//dane są wysyłane na kanał

//przyklad
var wg = sync.WaitGroup{}			//bez WaitGroup{} gorutyny nie zdążą się wykonać
func main(){
	ch := make(chan int)			//stworzenie dwukierunkowego kanału
	wg.Add(2)				
	go func(channel <-chan int){	//funckja jako argument przyjmuje receive only channel
		i := <-channel				//przyjmuje informacje z kanału
		fmt.Println(i)		
		wg.Done()
	}(ch)							//funkcja tylko odczytuje z kanału, ale może jako argument przyjąć kanał dwustronny
	go func(channel chan<- int){	//funkcja jako argument rzyjmuje send only channel
		ch <- 42					//funkcja wysyła informacje na kanał
		wg.Done()
	}(ch)							//funkcja tylko odczytuje z kanału, ale może jako argument przyjąć kanał dwustronny
	wg.Wait()
}

//buffered channel
channel := make(chan int, 50)		//channel może przechowywać w sobie 50 róznych intów

//używane głównie gdy chenele wysyłające i odbierające się nie równoważą
//zamykanie channeli i iterowanie po nich
var wg = sync.WaitGroup{}
func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(channel <-chan int) {
		for i := range channel {	//range leci po channelu i zwraca tylko jedną wartość pod danym miejscem
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(channel chan<- int) {
		ch <- 42
		ch <- 27
		close(channel)  //channel ma 50 miejsc, ale by nie było deadlocka to trzeba go zamknąć zanim druga gorutyna przeiteruje bo nieistniejących wartościach
		wg.Done()
	}(ch)
	wg.Wait()
}

//iterowanie po chanelu "nieskończoną" petlą
for {
	if i, ok := <- channel; ok {		//chennel zwraca dwie zmienne, wartość oraz bolleana ok, jeżeli jest ok to wydrukuje, w przeciwnym wypadku break
		fmt.Println(i)
	} else {
		break
	}
} 

//SELECT
//pozwala gorutynie namonitorowanie wielu channelow na raz
//przydatne w zamykaniu gorutyn poprzez przyjecie pustego structa{} do osobnego channela
var wg = sync.WaitGroup{}
var daneChan = make(chan string, 50)
var zrobioneChan = make(chan struct{}) //przyjmuje pustego structa (pusty struct w go nie zajmuje pamięci), tylko jako sygnał
func drukarz() {
	for {
		select { //sprawi, że zamknie oba kanały jesli nastapi break
		case entry := <-daneChan:
			fmt.Println(entry)
		case <-zrobioneChan:
			break
		}
	}
}
func main() {
	go drukarz()
	daneChan <- "costam costam"
	daneChan <- "elo elo"
	zrobioneChan <- struct{}{} //zainicjowanie pustego structa
} //bezpieczne zamkniecie gorutyny i kanałów