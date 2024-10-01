package Hangman

import (
	"fmt"
	"os"
)

//Quitter le jeu
func Quitter() {
	fmt.Println(" ")
	fmt.Println("Au revoir !")
	fmt.Println(" ")
	os.Exit(0)
}
