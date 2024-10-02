package Hangman

import (
	"math/rand"
	"time"
)
//Affichage de lettres en fonction de la taille du mot
func Help() {
	if len(mot) <= 5 {
		rand.Seed(time.Now().UnixNano())
		lmot := len(runes)
		indice1 := rand.Intn(lmot)
		lettre1 := runes[indice1]
		motstock[indice1] = string(runes[indice1])
		for i, lettre := range runes {
			if lettre == rune(lettre1) {
				motstock[i] = string(lettre)
			}
		}
	} else if len(mot) <= 8 && len(mot) > 5 {
		rand.Seed(time.Now().UnixNano())
		lmot := len(runes)
		indice1 := rand.Intn(lmot)
		indice2 := rand.Intn(lmot)
		lettre1 := runes[indice1]
		lettre2 := runes[indice2]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmot)
		}
		for i, lettre := range runes {
			if lettre == rune(lettre1) || lettre == rune(lettre2) {
				motstock[i] = string(lettre)
			}
		}
		motstock[indice1] = string(runes[indice1])
		motstock[indice2] = string(runes[indice2])
	} else if len(mot) > 8 {
		rand.Seed(time.Now().UnixNano())
		lmot := len(runes)
		indice1 := rand.Intn(lmot)
		indice2 := rand.Intn(lmot)
		indice3 := rand.Intn(lmot)
		lettre1 := runes[indice1]
		lettre2 := runes[indice2]
		lettre3 := runes[indice3]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmot)
		}
		for indice3 == indice2 || indice3 == indice1 {
			indice3 = rand.Intn(lmot)
		}
		for i, lettre := range runes {
			if lettre == lettre1 || lettre == rune(lettre2) || lettre == rune(lettre3) {
				motstock[i] = string(lettre)
			}
		}
		motstock[indice1] = string(runes[indice1])
		motstock[indice2] = string(runes[indice2])
		motstock[indice3] = string(runes[indice3])
	}
}
