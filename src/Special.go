package Hangman

import (
	"fmt"
	"time"
)

func Special() {
	ClearTerminal()
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║\033[30m\033[47m                 Listes de mots spéciaux                    \033[0m║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1 ► \033[95mMots Longs\033[0m                                             ║")
	fmt.Println("║                                                            ║")
	fmt.Println("║ 2 ► \033[95mMédecine\033[0m                                               ║")
	fmt.Println("║                                                            ║")
	fmt.Println("║ 3 ► \033[95mMécanique\033[0m                                              ║")
	fmt.Println("║                                                            ║")
	fmt.Println("║ 4 ► \033[95mInformatique\033[0m                                           ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║\033[30m\033[47m          Pour revenir au menu principal entrez 0           \033[0m║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println(" ")
	var input1 string
	fmt.Scanln(&input1)
	switch input1 {
	case "1":
		motslongs = true
		Pendu()
	case "2":
		medecine = true
		Pendu()
	case "3":

	case "4":

	case "0":
		MenuHangman()
		return
	default:
		fmt.Println(" ")
		fmt.Println("\033[91m\033[1mAucune fonction de mon jeu ne correspond à votre demande.\033[0m")
		fmt.Println(" ")
		time.Sleep(1500 * time.Millisecond)
		Special()
		return
	}
}
