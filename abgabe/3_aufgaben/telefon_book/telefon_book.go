package main

import (
	"fmt"
	"time"
)

var telefonBook = make(map[string]string)

// START OMIT
// run program
// END OMIT

func main() {
	/* Einträge */

	addTelefonBookEntries()

	/* Telefon ausgeben */
	fmt.Println("\nTelefon Nr. ausgeben")
	fmt.Println("Alan: Telefon #", telefonBook["alan"])
	fmt.Println("Philipp: Telefon #", telefonBook["pips"], "\n")
	<-time.After(time.Millisecond * 3000)

	/* Alan wechselt die Telefon # */
	fmt.Println("Alan wechselt Telefon #")
	changeTelefonBookEntry()
	<-time.After(time.Millisecond * 3000)

	/* Telefon ausgeben */
	fmt.Println("Alan's neue Nummer")
	fmt.Println("Alan: Telefon #", telefonBook["alan"], "\n")
	<-time.After(time.Millisecond * 3000)

	/* Philipp wirft Händy weg */
	fmt.Println("Philipp wirft händy weg!")
	deleteTelefonBookEntry()
	<-time.After(time.Millisecond * 3000)

	/* Telefon Philipp */
	fmt.Println("Philipp hat keine Nummer")
	fmt.Println("Philipp: Telefon #", telefonBook["pips"], "\n")
}

func addTelefonBookEntries() {
	telefonBook["alan"] = "+41 79 123 456"
	telefonBook["pips"] = "+41 79 456 123"
	telefonBook["dane"] = "+41 79 123 457"
}

func changeTelefonBookEntry() {
	telefonBook["alan"] = "+41 00 000 000"
}

func deleteTelefonBookEntry() {
	delete(telefonBook, "pips")
}
