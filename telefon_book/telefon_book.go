package main

import "fmt"

var telefonBook = make(map[string]string)

func main() {
	/* Einträge */
	telefonBook["alan"] = "+41 79 123 456"
	telefonBook["pips"] = "+41 79 456 123"
	telefonBook["dane"] = "+41 79 123 457"

	/* Telefon ausgeben */
	fmt.Println("\nTelefon Nr. ausgeben")
	fmt.Println("Alan: Telefon #", telefonBook["alan"])
	fmt.Println("Philipp: Telefon #", telefonBook["pips"], "\n")

	/* Alan wechselt die Telefon # */
	fmt.Println("Alan wechselt Telefon #")
	telefonBook["alan"] = "+41 00 000 000"

	/* Telefon ausgeben */
	fmt.Println("Alan's neue Nummer")
	fmt.Println("Alan: Telefon #", telefonBook["alan"], "\n")

	/* Philipp wirft Händy weg */
	fmt.Println("Philipp wirft händy weg!")
	delete(telefonBook, "pips")

	/* Telefon Philipp */
	fmt.Println("Philipp hat keine Nummer")
	fmt.Println("Philipp: Telefon #", telefonBook["pips"], "\n")
}
