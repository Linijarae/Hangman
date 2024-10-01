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

//Mode de jeu Difficile avec un mot de la liste Difficile
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
	dejafait := false
	for !strings.Contains(strings.Join(motstockD, ""), motD) {
		dejafait = false
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
		greeninput := "\033[92m" + input + "\033[0m"
		redinput := "\033[91m" + input + "\033[0m"
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
		if motD == input {
			ReussiteD()
			return
		} else if motD != input && len(input) > 1 {
			fmt.Println(" ")
			fmt.Println("\033[91mCe n'est pas le bon mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			erreur += 2
			inputstockmot = append(inputstockmot, redinput)
		}
		for j, r := range runesD {
			if string(r) == input && !dejafait {
				motstockD[j] = input
				correcte = true
				inputstock = append(inputstock, greeninput)
			}
		}
		if !correcte && len(input) == 1 && !dejafait {
			fmt.Println(" ")
			fmt.Println("\033[91mLa lettre ne fait pas partie du mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			inputstock = append(inputstock, redinput)
			erreur++
		}
	}
ReussiteD()
}
