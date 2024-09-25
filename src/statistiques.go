package Hangman

import "fmt"



func Statistiques() {
	fmt.Println(" ")
	fmt.Println("\033[4m\033[30m\033[47mVos Statistiques :\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Nombre d'innocents que vous n'avez pas réussi à sauver en facile :", "\033[91m",killF,"\033[0m")
	fmt.Println(" ")
	fmt.Println("Votre score actuel en facile :", "\033[92m",pointsF,"\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Nombre d'innocents que vous n'avez pas réussi à sauver en moyen :", "\033[91m",killM,"\033[0m")
	fmt.Println(" ")
	fmt.Println("Votre score actuel en moyen :", "\033[92m",pointsM,"\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Nombre d'innocents que vous n'avez pas réussi à sauver en difficile :", "\033[91m",killD,"\033[0m")
	fmt.Println(" ")
	fmt.Println("Votre score actuel en difficile :", "\033[92m",pointsD,"\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\033[30m\033[47m Pour revenir au menu principal entrez 0 \033[0m")
	fmt.Println(" ")
	var input1 string
	fmt.Scanln(&input1)
	switch input1 {
	case "0":
		fmt.Println(" ")
		MenuHangman()
		return
	default:
		fmt.Println(" ")
		fmt.Println("Aucune fonction de mon jeu ne correspond à votre demande.")
		fmt.Println(" ")
		Statistiques()
		return
	}
}
