package Hangman

import (
	"fmt"
	"time"
)

// Debut du jeu
func Start() {
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                               █    ├─┼─┤   ║")
	fmt.Println("║      ┌▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄         █    └─┴─┘   ║")
	fmt.Println("║      ├█┼┼┴┘  ║    ║  ╔╩╗  ╔╩╗  ║   ╔╩╗        █‗‗‗‗‗‗‗‗‗‗‗‗║")
	fmt.Println("║      ├█┼┘   ╔╩╗   ║  ║\033[91m\033[1mN\033[0m║  ║\033[91m\033[1mG\033[0m║  ║   ║\033[91m\033[1mD\033[0m║        █▄▄▄▄▄▄▄▄▄▄▄▄║")
	fmt.Println("║      ├█┤    ║\033[91m\033[1mH\033[0m║  ╔╩╗ ╚═╝  ╚═╝ ╔╩╗  ╚═╝        █   ┌────┐   ║")
	fmt.Println("║      ├█┤    ╚═╝  ║\033[91m\033[1mA\033[0m║          ║\033[91m\033[1mE\033[0m║             █   │    │   ║")
	fmt.Println("║      ├█┤         ╚═╝          ╚═╝             █   │    │   ║")
	fmt.Println("║      ├█┤                                      █   │¬   │   ║")
	fmt.Println("║      ├█┤                                      █   │    │   ║")
	fmt.Println("║   ▄▄▄███▄▄▄                                   █‗‗‗│    │‗‗‗║")
	fmt.Println("║▄████████████▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄██████████████║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║               ► ► ►  PRESS 1 TO START  ◄ ◄ ◄               ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println(" ")
	time.Sleep(2000 * time.Millisecond)
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("\033[1m\033[93mVous avez été enlevé, vous êtes attaché sur une chaise, baîlloné.\033[0m")
		fmt.Println(" ")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println("\033[95m\033[1m??????\033[0m - Bonjour, je vois que vous êtes réveillé. Nous allons jouer à un jeu. ")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println("\033[95m\033[1m??????\033[0m - Un pendu pour être plus précis. Mais ce pendu met dans la balance de vraie personnes.")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println("\033[95m\033[1m??????\033[0m - Dépassez les 150 points et je vous laisserai vivre. Et peut être même partir.")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println("\033[95m\033[1m??????\033[0m - Vous gagnerez des points par réussite, mais vous en perdrez pour chaque erreur.")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println("\033[95m\033[1m??????\033[0m - Ne devinez pas le mot, et ... vous savez ce qui se passera pour le pendu.")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println("\033[95m\033[1m??????\033[0m - Ne me décevez pas.")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[1m\033[93mIl enlève votre baillon, sort de la pièce, et vous laisse seul devant un vieil écran.\033[0m")
		fmt.Println(" ")
		time.Sleep(850 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[47m\033[30m                     Appuyez sur 0 pour continuer                    \033[0m")
		var input1 string
		fmt.Scanln(&input1)
		switch input1 {
		case "0":
			MenuHangman()
		default:
			fmt.Println("\033[91m\033[1mVous n'avez pas l'air de savoir suivre des instructions ... Bon courage pour la suite ...\033[0m")
			time.Sleep(2000 * time.Millisecond)
			MenuHangman()
		}
	default: 
	fmt.Println(" ")
		fmt.Println("\033[91m\033[1mAucune fonction de mon jeu ne correspond à votre demande.\033[0m")
		fmt.Println(" ")
		time.Sleep(1500 * time.Millisecond)
		Start()
		return
	}

}
