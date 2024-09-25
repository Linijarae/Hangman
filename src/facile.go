package Hangman

import (
	"fmt"
	"strings"
)

var killF = 0
var pointsF = 0
var facile = false
var lmotF = len(motF)
var motstockF = make([]string, lmotF)
var runesF = []rune(motF)

func Facile() {
	facile = true
	UnderscoreF()
	Erreur()
	for !strings.Contains(strings.Join(motstockF, ""), motF) {
		fmt.Println("Le mot actuel : ", strings.Join(motstockF, " "))
		fmt.Printf("Votre lettre : ")
		fmt.Scanln(&input)
		correcte := false
		for j, r := range runesF {
			if string(r) == input {
				motstockF[j] = input
				correcte = true
			}
		}
		if !correcte {
			fmt.Println("La lettre ne fait pas partie du mot !")
			erreur++
			Erreur()

		}
	}
	fmt.Println("Félicitations ! Vous avez deviné le mot :", strings.Join(motstockF, ""))
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m10\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsF = (pointsF + 10) - erreur
	fmt.Println("Vous obtenez :", "\033[92m", pointsF, "\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	erreur = 0
	facile = false
	if pointsF >= limitescore {
		Ending()
	}
	MenuHangman()
}
