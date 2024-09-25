package Hangman

import (
	"fmt"
	"time"
)

func Loading() {
	for t := 0; t < 5; t++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("Chargement en cours")
		fmt.Println(" ")
		fmt.Println(" . ")
		time.Sleep(200 * time.Millisecond)
		ClearTerminal()
		fmt.Println(" ")
		fmt.Println("Chargement en cours")
		fmt.Println(" ")
		fmt.Println(" .. ")
		time.Sleep(200 * time.Millisecond)
		ClearTerminal()
		fmt.Println(" ")
		fmt.Println("Chargement en cours")
		fmt.Println(" ")
		fmt.Println(" ... ")
		time.Sleep(200 * time.Millisecond)
		ClearTerminal()
	}
	Start()
}
