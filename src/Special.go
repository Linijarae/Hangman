package Hangman

import "fmt"

func Special() {
	fmt.Println(" ")
	fmt.Println("\033[4m\033[30m\033[47m ============= LISTES DE MOTS SPECIAUX ============= \033[0m")
	fmt.Println(" ")
	fmt.Println("1.- Mots Longs")
	fmt.Println(" ")
	fmt.Println("2.- Médecine")
	fmt.Println(" ")
	fmt.Println("3.- Mécanique")
	fmt.Println(" ")
	fmt.Println("4.- Informatique")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\033[4m\033[30m\033[47m              0.- Quitter              \033[0m")
	fmt.Println(" ")
	var input1 string
	fmt.Scanln(&input1)
	switch input1 {
	case "1":
		motslongs = true
		Pendu()
	case "2":

	case "3":
		
	case "4":
		
	case "0":
		Quitter()
	default:
		fmt.Println(" ")
		fmt.Println("\033[91m\033[1mAucune fonction de mon jeu ne correspond à votre demande.\033[0m")
		fmt.Println(" ")
		Special()
		return
	}
}