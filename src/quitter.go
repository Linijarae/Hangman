package Hangman

import (
	"fmt"
	"os"
)

//Quitter le jeu
func Quitter() {
	fmt.Println(" ")
	fmt.Println("Merci d'avoir joué ! Au plaisir de vous revoir !")
	fmt.Println(" ")
	os.Exit(0)
}
