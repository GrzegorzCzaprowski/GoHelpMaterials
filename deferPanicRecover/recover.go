package main

import (
	"fmt"
	"log"
)

func basicRecover() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}() //defer sprawia że nie wykonuje się fukcja bezpośrednio, tylko ja wzywa. te nawiasy są ważne
	panic("something bad happen")
	fmt.Println("end")
}

//recover jest przydatne tylko w zdeferowanych funkcjach, bo one jako jedyne wykonają się przed paniciem !!!
// obecna funkcja w której jest cover nie bedzie kontynuowana, ale wyzsze funkcje będą działac

/*
wydrukuje się:

start
2019/05/23 09:54:44 Error: something bad happen
+ jesli w mainie (poza funkcja basicRevocer) byłby dodatkowy kod to by się wykonal

recovery sprawia, że kod nie zostanie ubity, i main dalej się bedzie wykonywał ale to co jest w funkcji pod paniciem już sie nie wykona

*/

///////////

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func panickerRecover() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

/*
wydrukuje sie:

start
about to panic
2019/05/23 09:51:05 Error: something bad happened
end

cały program nie został ubity, zatrzymała się tylko funkcja, reszta programu wykonala się normalnie
jezeli usune defer, to funkcja sie spierdoli i caly program sie wysypie, panic "wygra". Gdy jest defer, to
zdeferowana funkcja wykona się przed paniciem
*/

//////////////

func repanicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func repanickerRecover() {
	fmt.Println("start")
	repanicker()
	fmt.Println("end")
}

/*wydrukuje się
start
about to panic
2019/05/23 10:18:05 Error: something bad happened
panic: something bad happened [recovered]
        panic: something bad happened

goroutine 1 [running]:
main.repanicker.func1()
        /home/grzegorz/go/zadanka/helpmaterials/recover.go:68 +0xb9
panic(0x4a2b80, 0x4dc410)
        /usr/lib/go/src/runtime/panic.go:522 +0x1b5
main.repanicker()
        /home/grzegorz/go/zadanka/helpmaterials/recover.go:71 +0xb2
main.repanickerRecover()
        /home/grzegorz/go/zadanka/helpmaterials/recover.go:77 +0x7f
main.main()
        /home/grzegorz/go/zadanka/helpmaterials/recover.go:82 +0x20
exit status 2
*/

func main() {
	repanickerRecover()

}
