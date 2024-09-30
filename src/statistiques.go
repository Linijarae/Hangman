package Hangman

import "fmt"



func Statistiques() {
	ClearTerminal()
	fmt.Println(" ")
	fmt.Println("\033[4m\033[30m\033[47m     Vos Statistiques :     \033[0m")
	fmt.Println(" ")
	fmt.Println("\033[4m\033[1mFACILE :\033[0m")
	fmt.Println("Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m",killF,"\033[0m")
	fmt.Println("Votre score :", "\033[92m",scoreF,"\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[4m\033[1mMOYEN :\033[0m")
	fmt.Println("Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m",killM,"\033[0m")
	fmt.Println("Votre score :", "\033[92m",scoreM,"\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[4m\033[1mDIFFICILE :\033[0m")
	fmt.Println("Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m",killD,"\033[0m")
	fmt.Println("Votre score :", "\033[92m",scoreD,"\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\033[30m\033[47m Pour revenir au menu principal entrez 0 \033[0m")
	fmt.Println(" ")
	var input1 string
	fmt.Scanln(&input1)
	switch input1 {
	case "0":
		ClearTerminal()
		MenuHangman()
		return
	default:
		fmt.Println(" ")
		fmt.Println("\033[91m\033[1mAucune fonction de mon jeu ne correspond à votre demande.\033[0m")
		fmt.Println(" ")
		Statistiques()
		return
	}
}
