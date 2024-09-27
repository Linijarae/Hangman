package Hangman

import (
	"fmt"
	"strings"
)

var killM = 0
var scoreM = 0
var moyen = false
var runesM []rune
var motstockM []string
var pointsM = 0

func Moyen() {
	MotM()
	moyen = true
	runesM = []rune(motM)
	lmotM := len(motM)
	motstockM = make([]string, lmotM)
	UnderscoreM()
	Erreur()
	HelpM()
	for !strings.Contains(strings.Join(motstockM, ""), motM) {
		fmt.Println("Le mot actuel : ", strings.Join(motstockM, " "))
		fmt.Printf("Votre lettre : ")
		fmt.Scanln(&input)
		correcte := false
		if string(runesM) == input {
			ReussiteM()
			return
		}
		for j, r := range runesM {
			if string(r) == input {
				motstockM[j] = input
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
ReussiteM()
}
