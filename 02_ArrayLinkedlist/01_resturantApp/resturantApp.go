package main

import (
	"fmt"
	"container/list"
)

//funzione che prepara il prossimo ordine della lista e chiama la comanda
func piattoPronto(l *list.List) {
	element := l.Front()
	l.Remove(element)
	 fmt.Println("ho preparato", element.Value)
}

func main() {
	//creo una nuova linked list 
	ordination := list.New()
	
	fmt.Println("Inserisci cosa desideri mangiare oggi: ")
	fmt.Println("Inserisci: \"fine\" per completare l'ordine")
	for {
		var item string
		fmt.Scan(&item)
		if item  == "fine" {
			break
		}
		ordination.PushFront(item)
	}

	fmt.Println()
	fmt.Println("La tua ordinazione comprende: ")
	for element := ordination.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
	fmt.Print("\n\n")
	
	
	//sezionedi codice dedicata alla preparazione dell'ordine
	fmt.Println("COMANDA:")
	for element := ordination.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for {
		//preparoo il proossimo  piatto della lista
		fmt.Println("sto preparando il piatto...")
		piattoPronto(ordination)
		
		if ordination.Front() == nil {
			break
		}
		fmt.Println()
		
		//piatti rimanenti
		fmt.Println("Piatti che rimangono da cucinare:")
		for element := ordination.Front(); element != nil; element = element.Next() {
			fmt.Println(element.Value)
		}	
	}

	fmt.Println("\nFine preparazione.")
}