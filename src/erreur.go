package Hangman

import (
	"fmt"
	"time"
)

var erreur int

// Progression du pendu en fonction du nombre d'erreurs
func Erreur() {
	if erreur == 0 {
		fmt.Println(" ")
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
		fmt.Println("======================================")
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
	} else if erreur >= 10 {
		fmt.Println(" ")
		fmt.Println("======================================")
		fmt.Println("\033[91m ")
		fmt.Println("     ________   ")
		fmt.Println("    |/      |   ")
		fmt.Println("    |      (_)  ")
		fmt.Println(`    |      /|\  `)
		fmt.Println("    |       |   ")
		fmt.Println(`    |      / \  `)
		fmt.Println("    |           ")
		fmt.Println("____|____       \033[0m")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("\033[1m\033[91mVous n'avez malheureusement pas réussi à sauver ce pauvre homme !\033[0m")
		time.Sleep(1500 * time.Millisecond)
		if facile {
			killF++
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mLe nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m", killF, "\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("\033[1mVotre score retombe de", "\033[92m", scoreF, "\033[0m", "\033[1mà 0 !\033[0m")
			fmt.Println(" ")
			fmt.Println("\033[1mLe mot était :", mot, "\033[0m")
			fmt.Println(" ")
			scoreF = 0
			facile = false
		}
		if moyen {
			killM++
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mLe nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m", killM, "\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("\033[1mVotre score retombe de", "\033[92m", scoreM, "\033[0m", "\033[1mà 0 !\033[0m")
			fmt.Println(" ")
			fmt.Println("\033[1mLe mot était :", mot, "\033[0m")
			fmt.Println(" ")
			scoreM = 0
			moyen = false
		}
		if difficile {
			killD++
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mLe nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m", killD, "\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("\033[1mVotre score retombe de", "\033[92m", scoreD, "\033[0m", "\033[1mà 0 !\033[0m")
			fmt.Println(" ")
			fmt.Println("\033[1mLe mot était :", mot, "\033[0m")
			fmt.Println(" ")
			scoreD = 0
			difficile = false
		}
		if motslongs {
			killL++
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mLe nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m", killL, "\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("\033[1mVotre score retombe de", "\033[92m", scoreL, "\033[0m", "\033[1mà 0 !\033[0m")
			fmt.Println(" ")
			fmt.Println("\033[1mLe mot était :", mot, "\033[0m")
			fmt.Println(" ")
			scoreL = 0
			motslongs = false
		}
		if medecine {
			killmed++
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mLe nombre de personnes que vous n'avez pas réussi à sauver s'élève désormais à :", "\033[91m", killmed, "\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("\033[1mVotre score retombe de", "\033[92m", scoremed, "\033[0m", "\033[1mà 0 !\033[0m")
			fmt.Println(" ")
			fmt.Println("\033[1mLe mot était :", mot, "\033[0m")
			fmt.Println(" ")
			scoremed = 0
			medecine = false
		}
		fmt.Println("\033[1mRetour au menu principal\033[0m")
		fmt.Println(" ")
		time.Sleep(5000 * time.Millisecond)
		erreur = 0
		ClearTerminal()
		MenuHangman()
	}
}
