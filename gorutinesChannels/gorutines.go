package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}

///////////////

func hello() {
	go sayHello()
}

//nic się nie wydrukuje, func main jest wykonywana od razy i gorutyna nie ma czsu by sie wykonac

///////////////////

func hello2() {
	go sayHello()
	fmt.Println("Bye")
}

//gorutyna ma szanse sie wykonać, ale na 90% nie zdaży tego zrobić

/////////////////

func hello3() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

//tutaj gorutyna zdazy sie wykonać, bo program sobie na coś czeka

///////////////

func changingVariableAfterGorutine() {
	text := "Hello"
	go func() {
		fmt.Println(text)
	}()
	text = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

//wydrukuje się "Goodbye", bo gorutyna czeka i wykonuje się dopiero po zmianie wartości text na goodbye. funkcja nie wykonuje sie od razu
//nie jest to gwarantowane, jest szansa, że wydrukuje się Hello

/////////////////////

func changingVariableAfterGorutine2() {
	text := "Hello"
	go func(text string) {
		fmt.Println(text)
	}(text)
	text = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

//w 100% przypadków wydrukuje się Hello, bo argument do gorutyny został przekazany wczesniej, nawet jesli gorutyna wykonała się po zmianie zmiennej na inną wartość

func main() {
	changingVariableAfterGorutine2()
}
