package Hangman

import (
	"fmt"
	"time"
)

var killF = 0
var scoreF = 0
var pointsF int
var killM = 0
var scoreM = 0
var pointsM int
var killD = 0
var scoreD = 0
var pointsD int

// Statistiques des difficultés facile, moyen et difficile
func Statistiques() {
	ClearTerminal()
	fmt.Println("╔═════════════════════════════════════════════════════════════════╗")
	fmt.Println("║\033[30m\033[47m                           Statistiques                          \033[0m║")
	fmt.Println("╚═════════════════════════════════════════════════════════════════╝")
	fmt.Println(" ")
	fmt.Println("\033[1m\033[92m ► \033[4mFACILE :\033[0m")
	fmt.Println("    Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m", killF, "\033[0m")
	fmt.Println("    Votre score :", "\033[92m", scoreF, "\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[1m\033[93m ► \033[4mMOYEN :\033[0m")
	fmt.Println("    Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m", killM, "\033[0m")
	fmt.Println("    Votre score :", "\033[92m", scoreM, "\033[0m")
	fmt.Println(" ")
	fmt.Println("\033[1m\033[91m ► \033[4mDIFFICILE :\033[0m")
	fmt.Println("    Nombre d'innocents que vous n'avez pas réussi à sauver :", "\033[91m", killD, "\033[0m")
	fmt.Println("    Votre score :", "\033[92m", scoreD, "\033[0m")
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
	case "2":
		StatistiquesSpe()
	case "0":
		MenuHangman()
		return
	default:
		fmt.Println(" ")
		fmt.Println("\033[91m\033[1mAucune fonction de mon jeu ne correspond à votre demande.\033[0m")
		fmt.Println(" ")
		time.Sleep(1500 * time.Millisecond)
		Statistiques()
		return
	}
}
