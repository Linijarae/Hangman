package Hangman

import (
	"fmt"
	"strings"
)

var killD = 0
var pointsD = 0
var difficile = false
var runesD []rune
var motstockD []string
var scoreD = 0


func Difficile() {
	difficile = true
	MotD()
	runesD = []rune(motD)
	lmotD := len(motD)
	motstockD = make([]string, lmotD)
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
			fmt.Println(" ")
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			fmt.Println(" ")
			erreur++
			Erreur()

		}
	}
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :", strings.Join(motstockD, ""))
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m15\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsD = 15 - erreur
	scoreD += pointsD
	fmt.Println("vous obtenez :","\033[92m", pointsD, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :", "\033[92m", scoreD, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println(" ")
	difficile = false
	erreur = 0
	pointsD = 0
	if scoreD >= limitescore {
		Ending()
	}
	MenuHangman()
}
