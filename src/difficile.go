package Hangman

import (
	"fmt"
	"strings"
)

var killD = 0
var pointsD = 0
var difficile = false
var lmotD = len(motD)
var motstockD = make([]string, lmotD)
var runesD = []rune(motD)

func Difficile() {
	difficile = true
	MotD()
	UnderscoreD()
	Erreur()
	for !strings.Contains(strings.Join(motstockD, ""), motD) {
		fmt.Println("Le mot actuel : ", strings.Join(motstockD, " "))
		fmt.Printf("Votre lettre : ")
		fmt.Scanln(&input)
		correcte := false
		for j, r := range runesD {
			if string(r) == input {
				motstockD[j] = input
				correcte = true
			}
		}
		if !correcte {
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			erreur++
			Erreur()

		}
	}
	fmt.Println("Félicitations ! Vous avez deviné le mot :", strings.Join(motstockD, ""))
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m15\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsD = (pointsD + 15) - erreur
	fmt.Println("Vous obtenez :", "\033[92m", pointsD, "\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	difficile = false
	erreur = 0
	if pointsD >= limitescore {
		Ending()
	}
	MenuHangman()
}
