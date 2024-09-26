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
	for !strings.Contains(strings.Join(motstockM, ""), motM) {
		fmt.Println("Le mot actuel : ", strings.Join(motstockM, " "))
		fmt.Printf("Votre lettre : ")
		fmt.Scanln(&input)
		correcte := false
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
	fmt.Println("Félicitations ! Vous avez deviné le mot :", strings.Join(motstockM, ""))
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m12\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsM = 12 - erreur
	scoreM += pointsM
	fmt.Println("vous obtenez :","\033[92m",pointsM,"\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :","\033[92m",scoreM, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println(" ")
	erreur = 0
	moyen = false
	pointsM = 0
	if scoreM >= limitescore {
		Ending()
	}
	MenuHangman()
}
