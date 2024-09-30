package Hangman

import (
	"fmt"
	"strings"
	"time"
)

var killD = 0
var pointsD = 0
var difficile = false
var runesD []rune
var motstockD []string
var scoreD = 0

func Difficile() {
	MotD()
	difficile = true
	runesD = []rune(motD)
	lmotD := len(motD)
	motstockD = make([]string, lmotD)
	var inputstock []string
	var inputstockmot []string
	UnderscoreD()
	HelpD()
	for !strings.Contains(strings.Join(motstockD, ""), motD) {
		ClearTerminal()
		Erreur()
		fmt.Println(" ")
		fmt.Println("\033[4m\033[93mLe mot actuel :\033[0m", strings.Join(motstockD, " "))
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
		if motD == input {
			ReussiteD()
			return
		} else if motD != input && len(input) > 1 {
			fmt.Println(" ")
			fmt.Println("\033[91mCe n'est pas le bon mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(2000 * time.Millisecond)
			erreur += 2
			redinput := "\033[91m" + input + "\033[0m"
			inputstockmot = append(inputstockmot, redinput)
		}
		for j, r := range runesD {
			if string(r) == input {
				motstockD[j] = input
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

}
