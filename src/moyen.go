package Hangman

import (
	"fmt"
	"strings"
)

var killM = 0
var pointsM = 0
var moyen = false
var lmotM = len(motM)
var motstockM = make([]string, lmotM)
var runesM = []rune(motM)

func Moyen() {
	moyen = true
	MotM()
	UnderscoreM()
	Erreur()
	for !strings.Contains(strings.Join(motstockM, ""), motM) {
		fmt.Println("Le mot actuel : ", strings.Join(motstockM, " "))
		fmt.Printf("Votre lettre : ")
		fmt.Scanln(&input)
		correcte := false
		for j, r := range runesM {
			if string(r) == input {
				motstockM[j] = input
				correcte = true
			}
		}
		if !correcte {
			fmt.Println("La lettre ne fait pas partie du mot !")
			erreur++
			Erreur()

		}
	}
	fmt.Println("Félicitations ! Vous avez deviné le mot :", strings.Join(motstockM, ""))
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m12\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsM = (pointsM + 12) - erreur
	fmt.Println("Vous obtenez :", "\033[92m", pointsM, "\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	erreur = 0
	moyen = false
	if pointsM >= limitescore {
		Ending()
	}
	MenuHangman()
}
