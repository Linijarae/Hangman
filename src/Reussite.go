package Hangman

import (
	"fmt"
	"time"
)

// Mot facile trouvé
func ReussiteF() {
	ClearTerminal()
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
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
	time.Sleep(5000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("Retour au menu principal.")
	fmt.Println(" ")
	time.Sleep(2000 * time.Millisecond)
	MenuHangman()
}

// Mot moyen trouvé
func ReussiteM() {
	ClearTerminal()
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
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
	time.Sleep(5000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("Retour au menu principal.")
	fmt.Println(" ")
	time.Sleep(2000 * time.Millisecond)
	MenuHangman()
}

// Mot difficile trouvé
func ReussiteD() {
	ClearTerminal()
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
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
	time.Sleep(5000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("Retour au menu principal.")
	fmt.Println(" ")
	time.Sleep(2000 * time.Millisecond)
	MenuHangman()
}

// Mot Long trouvé
func ReussiteL() {
	ClearTerminal()
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m25\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsL = 25 - erreur
	scoreL += pointsL
	fmt.Println("vous obtenez :", "\033[92m", pointsL, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :", "\033[92m", scoreL, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println(" ")
	motslongs = false
	erreur = 0
	pointsL = 0
	if scoreL >= limitescore {
		Ending()
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("Retour au menu principal.")
	fmt.Println(" ")
	time.Sleep(2000 * time.Millisecond)
	MenuHangman()
}

// Mot Medecine trouvé
func ReussiteMed() {
	ClearTerminal()
	fmt.Println(" ")
	fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
	fmt.Println(" ")
	fmt.Println("Vous obtenez \033[92m25\033[0m points")
	fmt.Println(" ")
	fmt.Println("Mais ... Auxquels je retire le nombre de vos erreurs ... et ...")
	fmt.Println(" ")
	pointsmed = 25 - erreur
	scoremed += pointsmed
	fmt.Println("vous obtenez :", "\033[92m", pointsmed, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println("Vous avez :", "\033[92m", scoremed, "\033[0mpoints !")
	fmt.Println(" ")
	fmt.Println(" ")
	medecine = false
	erreur = 0
	pointsmed = 0
	if scoremed >= limitescore {
		Ending()
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("Retour au menu principal.")
	fmt.Println(" ")
	time.Sleep(2000 * time.Millisecond)
	MenuHangman()
}
