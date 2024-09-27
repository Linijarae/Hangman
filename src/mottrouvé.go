package Hangman

import (
	"fmt"
)

func ReussiteF() {
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :",motF)
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m10\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsF = 10 - erreur
	scoreF += pointsF
	fmt.Println("vous obtenez :", "\033[92m", pointsF, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :", "\033[92m", scoreF, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println(" ")
	erreur = 0
	facile = false
	pointsF = 0
	if scoreF >= limitescore {
		Ending()
	}
	MenuHangman()
}

func ReussiteM() {
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :",motM)
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m12\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsM = 12 - erreur
	scoreM += pointsM
	fmt.Println("vous obtenez :", "\033[92m", pointsM, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :", "\033[92m", scoreM, "\033[0mpoints !")
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

func ReussiteD() {
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :", motD)
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m15\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsD = 15 - erreur
	scoreD += pointsD
	fmt.Println("vous obtenez :", "\033[92m", pointsD, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :", "\033[92m", scoreD, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println(" ")
	difficile = false
	erreur = 0
	pointsD = 0
	if scoreD >= limitescore {
		Ending()
	}
	MenuHangman()
}
