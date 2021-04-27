/*
Scrivere un programma 'rettangoli.go' che legge da standard input
righe che definiscono "rettangoli" (nome, alt, largh),
e crea un rettangolo per ogni riga incontrata.

Nel programma definire una struttura Rettangolo con campi:
- 'nome' di tipo string
- 'alt' e 'largh' di tipo float64

Definire inoltre le seguenti funzioni:

- 	func maggiore(r1, r2 Rettangolo) boolean
		che restituisce 'true' se il rettangolo r1 ha AREA maggiore di r2

- 	func allarga(r1 *Rettangolo, zoom float64)
		che ingrandisce r1 (moltiplica altezza e larghezza) di un fattore 'zoom'

Definire un main che per ogni linea letta:
- crea il rettangolo corrispondente;
- se la linea ha 4 dati, allarga il rettangolo stesso del fattore
fornito come quarto dato;
- trova il più grande rettangolo tra quelli letti da stdin (fine input ctrl^D)
e lo stampa.

Per esempio in corrispondenza della linea:

uno 3 2 10

il programma deve creare il rettangolo (uno,3,2)
e poi allargarlo di un fattore 10.

Esempio
-------
Nota: le righe di input sono contrassegnate da <- (che non andrà digitato)
per distinguerle dall'output generato dal programma

$ go run rettangoli.go
<-uno 3 2
<-due 6 7
<-tre 8 67 2
<-ctrl D
max {tre 16 134}

NB: il rettangolo "tre" è stato "zoomato" di un fattore 2

*/

package main

import (
	"testing"
)

func TestOK(t *testing.T) {
	// questi li accorpo, sono entrambi output OK
	lanciaGenerica(t, "./rettangoli", "uno 3 2\ndue 6 7\ntre 8 67 2\n", "max {tre 16 134}\n")

	lanciaGenerica(t, "./rettangoli", "primo 5 999\nsecondo 6 4\nterzo 31231 1\n", "max {terzo 31231 1}\n")
}

func TestAllargaZero(t *testing.T) {
	var rett Rettangolo
	allarga(&rett, 4)

	if rett.alt != 0 || rett.largh != 0 {
		t.Fail()
	}
}

func TestAllargaNonZeroEMaggiore(t *testing.T) {
	rett1 := Rettangolo{"uno", 5, 10}
	rett2 := Rettangolo{"due", 6, 11}

	if maggiore(rett1, rett2) {
		t.Fail()
	}

	allarga(&rett1, 4) // rett1 diventa area maggiore

	if !maggiore(rett1, rett2) {
		t.Fail()
	}
}

func TestStringeNonZeroEMaggiore(t *testing.T) {
	rett1 := Rettangolo{"uno", 5, 10}
	rett2 := Rettangolo{"due", 4, 10}

	if !maggiore(rett1, rett2) {
		t.Fail()
	}

	allarga(&rett1, .4) // rett1 diventa area minore

	if maggiore(rett1, rett2) {
		t.Fail()
	}
}
