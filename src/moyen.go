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
	var inputstock []string
	var inputstockmot []string
	UnderscoreM()
	Erreur()
	HelpM()
	for !strings.Contains(strings.Join(motstockM, ""), motM) {
		fmt.Println(" ")
		fmt.Println("Le mot actuel : ", strings.Join(motstockM, " "))
		fmt.Println(" ")
		fmt.Println("\nVous avez utilisé les lettres :", inputstock)
		fmt.Println(" ")
		fmt.Println("\nVous avez tenté les mots :", inputstockmot)
		fmt.Println(" ")
		fmt.Scanln(&input)
		correcte := false
		for _, a := range inputstock {
			if a == input {
				fmt.Println(" ")
				fmt.Println("\033[91mVous avez déjà essayé cette lettre.\033[0m")
				fmt.Println(" ")
			}
		}
		if string(runesM) == input {
			ReussiteM()
			return
		} else if motF != input && len(input) > 1 {
			fmt.Println(" ")
			fmt.Println("\033[91mCe n'est pas le bon mot !\033[0m")
			fmt.Println(" ")
			erreur += 2
			redinput := "\033[91m" + input + "\033[0m"
			inputstockmot = append(inputstockmot, redinput)
			Erreur()
		}
		for j, r := range runesM {
			if string(r) == input {
				motstockM[j] = input
				correcte = true
				greeninput := "\033[92m" + input + "\033[0m"
				inputstock = append(inputstock, greeninput)
			}
		}
		if !correcte {
			fmt.Println(" ")
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			fmt.Println(" ")
			redinput := "\033[91m" + input + "\033[0m"
			inputstock = append(inputstock, redinput)
			erreur++
			Erreur()

		}
	}
ReussiteM()
}
