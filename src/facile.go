package Hangman

import (
	"fmt"
	"strings"
)

var killF = 0
var scoreF = 0
var facile = false
var runesF []rune
var motstockF []string
var pointsF = 0

func Facile() {
	MotF()
	facile = true
	runesF = []rune(motF)
	lmotF := len(motF)
	motstockF = make([]string, lmotF)
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
			fmt.Println(" ")
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			fmt.Println(" ")
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
	pointsF = 10 - erreur
	scoreF += pointsF
	fmt.Println("vous obtenez :","\033[92m", pointsF, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :","\033[92m", scoreF, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println(" ")
	erreur = 0
	facile = false
	pointsF = 0
	if scoreF >= limitescore {
		Ending()
	}
	MenuHangman()
}
