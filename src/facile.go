package Hangman

import (
	"fmt"
	"strings"
	"time"
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
	var inputstockmot []string
	UnderscoreF()
	HelpF()
	for !strings.Contains(strings.Join(motstockF, ""), motF) {
		ClearTerminal()
		Erreur()
		fmt.Println(" ")
		fmt.Println("\033[4m\033[93mLe mot actuel :\033[0m", strings.Join(motstockF, " "))
		fmt.Println(" ")
		fmt.Println("\033[4m\033[93mVous avez utilisé les lettres :\033[0m", inputstock)
		fmt.Println(" ")
		fmt.Println("\033[4m\033[93mVous avez tenté les mots :\033[0m", inputstockmot)
		fmt.Println(" ")
		fmt.Scanln(&input)
		correcte := false
		for _, a := range inputstock {
			if a == input {
				fmt.Println(" ")
				fmt.Println("\033[91mVous avez déjà essayé cette lettre.\033[0m")
				fmt.Println(" ")
				time.Sleep(2000 * time.Millisecond)
				break
			}
		}
		if motF == input {
			ReussiteF()
			return
		} else if motF != input && len(input) > 1 {
			fmt.Println(" ")
			fmt.Println("\033[91mCe n'est pas le bon mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(2000 * time.Millisecond)
			erreur += 2
			redinput := "\033[91m" + input + "\033[0m"
			inputstockmot = append(inputstockmot, redinput)
		}
		for j, r := range runesF {
			if string(r) == input {
				motstockF[j] = input
				correcte = true
				greeninput := "\033[92m" + input + "\033[0m"
				inputstock = append(inputstock, greeninput)
			}
		}
		if !correcte && len(input) == 1 {
			fmt.Println(" ")
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(2000 * time.Millisecond)
			redinput := "\033[91m" + input + "\033[0m"
			inputstock = append(inputstock, redinput)
			erreur++
		}

	}
	ReussiteF()
}
