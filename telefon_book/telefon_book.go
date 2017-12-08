package main

import "fmt"

var telefonBook = make(map[string]string)

func main() {
	/* Einträge */
	addTelefonBookEntries()

	/* Telefon ausgeben */
	fmt.Println("\nTelefon Nr. ausgeben")
	fmt.Println("Alan: Telefon #", telefonBook["alan"])
	fmt.Println("Philipp: Telefon #", telefonBook["pips"], "\n")

	/* Alan wechselt die Telefon # */
	fmt.Println("Alan wechselt Telefon #")
	changeTelefonBookEntry()

	/* Telefon ausgeben */
	fmt.Println("Alan's neue Nummer")
	fmt.Println("Alan: Telefon #", telefonBook["alan"], "\n")

	/* Philipp wirft Händy weg */
	fmt.Println("Philipp wirft händy weg!")
	deleteTelefonBookEntry()

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
