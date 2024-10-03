package Hangman

import (
	"fmt"
	"time"
)

var killL = 0
var scoreL = 0
var pointsL int
var killmed = 0
var scoremed = 0
var pointsmed int

// Statistiques de chaque difficultées
func StatistiquesSpe() {
	ClearTerminal()
	fmt.Println("╔═════════════════════════════════════════════════════════════════╗")
	fmt.Println("║\033[30m\033[47m                     Statistiques en Spécial                     \033[0m║")
	fmt.Println("╚═════════════════════════════════════════════════════════════════╝")
	fmt.Println(" ")
	fmt.Println("\033[1m\033[95m ► \033[4mMOTS LONGS :\033[0m")
	fmt.Println("    Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m", killL, "\033[0m")
	fmt.Println("    Votre score :", "\033[92m", scoreL, "\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[1m\033[95m ► \033[4mMEDECINE :\033[0m")
	fmt.Println("    Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m", killmed, "\033[0m")
	fmt.Println("    Votre score :", "\033[92m", scoremed, "\033[0m")
	fmt.Println(" ")
	fmt.Println("╔═════════════════════════════════════════════════════════════════╗")
	fmt.Println("║      \033[36mEntrez le numéro de la page que vous souhaitez voir\033[0m        ║")
	fmt.Println("║\033[1m                         PAGE \033[107m\033[30m 1 \033[0m\033[1m -  2                           \033[0m║")
	fmt.Println("╠═════════════════════════════════════════════════════════════════╣")
	fmt.Println("║\033[30m\033[47m             Pour revenir au menu principal entrez 0             \033[0m║")
	fmt.Println("╚═════════════════════════════════════════════════════════════════╝")
	fmt.Println(" ")
	var input1 string
	fmt.Scanln(&input1)
	switch input1 {
	case "1":
		Statistiques()
		return
	case "2":
		StatistiquesSpe()
		return
	case "0":
		MenuHangman()
		return
	default:
		fmt.Println(" ")
		fmt.Println("\033[91m\033[1mAucune fonction de mon jeu ne correspond à votre demande.\033[0m")
		fmt.Println(" ")
		time.Sleep(1500 * time.Millisecond)
		StatistiquesSpe()
		return
	}
}
