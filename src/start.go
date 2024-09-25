package Hangman

import (
	"fmt"
	"time"
)

func Start() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("\033[1mVous avez été enlevé, vous êtes attaché sur une chaise, baîlloné.\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("?????? - Bonjour, je vois que vous êtes réveillé. Nous allons jouer à un jeu. ")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("?????? - Un pendu pour être plus précis. Mais ce pendu met dans la balance de vraie personnes.")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("?????? - Dépassez les 150 points et je vous laisserai vivre. Et peut être même partir.")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("?????? - Vous gagnerez des points par réussite, mais vous en perdrez pour chaque erreur.")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("?????? - Ne devinez pas le mot, et ... vous savez ce qui se passera pour le pendu.")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("?????? - Ne me décevez pas.")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("\033[1mIl enlève votre baillon, sort de la pièce, et vous laisse seul devant un vieil écran.\033[0m")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("\033[47m\033[30m          Appuyez sur 0 pour continuer         \033[0m")
	var input1 string
	fmt.Scanln(&input1)
	switch input1 {
	case "0":
		ClearTerminal()
		MenuHangman()
	default:
		fmt.Println("\033[1mVous n'avez pas l'air de savoir suivre des instructions ... Bon courage pour la suite ...\033[0m")
		time.Sleep(2000 * time.Millisecond)
		ClearTerminal()
	}
}
