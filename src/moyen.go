package Hangman

import (
	"fmt"
	"strings"
	"time"
)

var killM = 0
var scoreM = 0
var moyen = false
var runesM []rune
var motstockM []string
var pointsM = 0

//Mode de jeu Moyen avec un mot de la liste Moyen
func Moyen() {
	MotM()
	moyen = true
	runesM = []rune(motM)
	lmotM := len(motM)
	motstockM = make([]string, lmotM)
	var inputstock []string
	var inputstockmot []string
	UnderscoreM()
	HelpM()
	dejafait := false
	for !strings.Contains(strings.Join(motstockM, ""), motM) {
		dejafait = false
		ClearTerminal()
		Erreur()
		fmt.Println(" ")
		fmt.Println("\033[4m\033[93mLe mot actuel :\033[0m", strings.Join(motstockM, " "))
		fmt.Println(" ")
		fmt.Println("\033[4m\033[93mVous avez utilisé les lettres :\033[0m", inputstock)
		fmt.Println(" ")
		fmt.Println("\033[4m\033[93mVous avez tenté les mots :\033[0m", inputstockmot)
		fmt.Println(" ")
		fmt.Scanln(&input)
		correcte := false
		redinput := "\033[91m" + input + "\033[0m"
		greeninput := "\033[92m" + input + "\033[0m"
		for _, a := range inputstock {
			if a == redinput || a == greeninput {
				fmt.Println(" ")
				fmt.Println("\033[91mVous avez déjà essayé cette lettre.\033[0m")
				fmt.Println(" ")
				dejafait = true
				time.Sleep(1500 * time.Millisecond)
				break
			}
		}
		if motM == input {
			ReussiteM()
			return
		} else if motM != input && len(input) > 1 {
			fmt.Println(" ")
			fmt.Println("\033[91mCe n'est pas le bon mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			erreur += 2
			inputstockmot = append(inputstockmot, redinput)
		}
		for j, r := range runesM {
			if string(r) == input && !dejafait {
				motstockM[j] = input
				correcte = true
				inputstock = append(inputstock, greeninput)
			}
		}
		if !correcte && len(input) == 1 && !dejafait{
			fmt.Println(" ")
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			inputstock = append(inputstock, redinput)
			erreur++
		}
	}
	ReussiteM()
}
