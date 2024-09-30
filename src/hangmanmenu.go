package Hangman

import "fmt"

var input string
var limitescore = 150

func MenuHangman() {
	fmt.Println(" ")
	fmt.Println("\033[4m\033[30m\033[47m      Séléctionnez votre niveau :      \033[0m")
	fmt.Println(" ")
	fmt.Println("1.- \033[92mFacile\033[0m")
	fmt.Println(" ")
	fmt.Println("2.- \033[93mMoyen\033[0m")
	fmt.Println(" ")
	fmt.Println("3.- \033[91mDifficile\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("4.- \033[96mVos Statistiques \033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\033[4m\033[30m\033[47m              0.- Quitter              \033[0m")
	fmt.Println(" ")
	var input1 string
	fmt.Scanln(&input1)
	switch input1 {
	case "1":
		Facile()
	case "2":
		Moyen()
	case "3":
		Difficile()
	case "4":
		Statistiques()
	case "0":
		Quitter()
	default:
		fmt.Println(" ")
		fmt.Println("\033[91m\033[1mAucune fonction de mon jeu ne correspond à votre demande.\033[0m")
		fmt.Println(" ")
		MenuHangman()
		return
	}
}
