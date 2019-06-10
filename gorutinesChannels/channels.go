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