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
	HelpD()
	for !strings.Contains(strings.Join(motstockD, ""), motD) {
		fmt.Println("Le mot actuel : ", strings.Join(motstockD, " "))
		fmt.Printf("Votre lettre : ")
		fmt.Scanln(&input)
		correcte := false
		if string(runesD) == input {
			ReussiteD()
			return
		}
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

}
