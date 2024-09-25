package Hangman

import (
	"fmt"
	"time"
)

var erreur int

func Erreur() {
	if erreur == 0 {
		fmt.Println(" ")
		fmt.Println("                ")
		fmt.Println("                ")
		fmt.Println("                ")
		fmt.Println(`                `)
		fmt.Println("                ")
		fmt.Println(`                `)
		fmt.Println("                ")
		fmt.Println("_________       ")
		fmt.Println(" ")
	} else if erreur == 1 {
		fmt.Println(" ")
		fmt.Println("                ")
		fmt.Println("    |           ")
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 2 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |           ")
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 3 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/          ")
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 4 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 5 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |      (_)  ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 6 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |      (_)  ")
		fmt.Println(`    |       |   `)
		fmt.Println("    |       |   ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 7 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |      (_)  ")
		fmt.Println(`    |      /|   `)
		fmt.Println("    |       |   ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 8 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |      (_)  ")
		fmt.Println(`    |      /|\  `)
		fmt.Println("    |       |   ")
		fmt.Println(`    |           `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 9 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |      (_)  ")
		fmt.Println(`    |      /|\  `)
		fmt.Println("    |       |   ")
		fmt.Println(`    |      /    `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
	} else if erreur == 10 {
		fmt.Println(" ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |      (_)  ")
		fmt.Println(`    |      /|\  `)
		fmt.Println("    |       |   ")
		fmt.Println(`    |      / \  `)
		fmt.Println("    |           ")
		fmt.Println("____|____       ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Vous n'avez malheureusement pas réussi à sauver ce pauvre homme !")
		time.Sleep(1500 * time.Millisecond)
		if facile {
			killF++
		time.Sleep(1500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("Le nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m",killF,"\033[0m")
		fmt.Println(" ")
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("Votre score retombe de","\033[92m",pointsF,"\033[0m","à 0 !")
		fmt.Println(" ")
			pointsF = 0
			facile = false
		}
		if moyen {
			killM++
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("Le nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m",killM,"\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("Votre score retombe de","\033[92m",pointsM,"\033[0m","à 0 !")
			fmt.Println(" ")
			pointsM = 0
			moyen = false
		}
		if difficile {
			killD++
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("Le nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m",killD,"\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("Votre score retombe de","\033[92m",pointsD,"\033[0m","à 0 !")
			fmt.Println(" ")
			pointsD = 0
			difficile = false
		}
		fmt.Println("Retour au menu principal")
		fmt.Println(" ")
		time.Sleep(2000 * time.Millisecond)
		erreur = 0
		MenuHangman()
	}
}
