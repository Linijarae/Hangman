package Hangman

import "fmt"

func MenuHangman(){
	fmt.Println("Bienvenue dans mon jeu du pendu !")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Séléctionnez votre niveau :")
	fmt.Println(" ")
	fmt.Println("1.- Facile")
	fmt.Println(" ")
	fmt.Println("2.- Moyen")
	fmt.Println(" ")
	fmt.Println("3.- Difficile")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("0.- Quitter")
	fmt.Println(" ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		
	case "2":

	case "3":

	case "0":
		Quitter()
	}
}