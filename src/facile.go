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
var pointsF int

func Facile() {
	MotF()
	facile = true
	runesF = []rune(motF)
	lmotF := len(motF)
	motstockF = make([]string, lmotF)
	var inputstock []string
	UnderscoreF()
	Erreur()
	HelpF()
	for !strings.Contains(strings.Join(motstockF, ""), motF) {
		fmt.Println(" ")
		fmt.Println("Le mot actuel : ", strings.Join(motstockF, " "))
		fmt.Println(" ")
		fmt.Println("\nVous avez utilisé les lettres :", inputstock)
		fmt.Println(" ")
		fmt.Scanln(&input)
		correcte := false
		for _, a := range inputstock {
			if a == input {
				fmt.Println("Vous avez déjà essayé cette lettre.")
			}
		}
		if string(runesF) == input {
			ReussiteF()
			return
		}
		for j, r := range runesF {
			if string(r) == input {
				motstockF[j] = input
				correcte = true
				inputstock = append(inputstock, input)
			}
		}
		if !correcte {
			fmt.Println(" ")
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			fmt.Println(" ")
			inputstock = append(inputstock, input)
			erreur++
			Erreur()

		}
	}
	ReussiteF()
}
