type mutex chan struct{} //dodajemy tą strukture jako pole do innej struktury, która chcemy by wykonywała sie niejednoczesnie

func (m mutex) Lock() {
	m <- struct{}{} //kanał mutex przyjmie każdą strukturę
}
func (m mutex) Unlock() {
	<-m //mutex zostaje zwolniony
}

type PrzykladowaStruktura struct {
	name  string
	email string
	mutex mutex
}

//zainicjowanie struktury:
przyklad := &PrzykladowaStruktura{
	name: "asdasdsa",
	email: "sadasdsad",
	mutex: make(chan struct{}, 1),
}

//metody dla danej struktury ktora korzysta z mutexa i nie wystepuje race:

func (p *PrzykladowaStrutura) ZrobCos(){
	p.mutex.Lock()
	defer p.mutex.Unlock()
	//tu się coś robi, reszta ciała metody
}