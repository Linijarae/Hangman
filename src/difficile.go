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
	var inputstock []string
	var inputstockmot []string
	UnderscoreD()
	Erreur()
	HelpD()
	for !strings.Contains(strings.Join(motstockD, ""), motD) {
		fmt.Println(" ")
		fmt.Println("Le mot actuel : ", strings.Join(motstockD, " "))
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
				fmt.Println("Vous avez déjà essayé cette lettre.")
				fmt.Println(" ")
			}
		}
		if string(runesD) == input {
			ReussiteD()
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
		for j, r := range runesD {
			if string(r) == input {
				motstockD[j] = input
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

}
