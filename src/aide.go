package Hangman

import (
	"math/rand"
	"time"
)

func HelpF() {
	if len(motF) <= 5 {
		rand.Seed(time.Now().UnixNano())
		lmotF := len(runesF)
		indice1 := rand.Intn(lmotF)
		lettre1 := runesF[indice1]
		motstockF[indice1] = string(runesF[indice1])
		for i, lettre := range runesF {
			if lettre == rune(lettre1) {
				motstockF[i] = string(lettre)
			}
		}
	} else if len(motF) <= 8 && len(motF) > 5 {
		rand.Seed(time.Now().UnixNano())
		lmotF := len(runesF)
		indice1 := rand.Intn(lmotF)
		indice2 := rand.Intn(lmotF)
		lettre1 := runesF[indice1]
		lettre2 := runesF[indice2]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmotF)
		}
		for i, lettre := range runesF {
			if lettre == rune(lettre1) || lettre == rune(lettre2) {
				motstockF[i] = string(lettre)
			}
		}
		motstockF[indice1] = string(runesF[indice1])
		motstockF[indice2] = string(runesF[indice2])
	} else if len(motF) < 8 {
		rand.Seed(time.Now().UnixNano())
		lmotF := len(runesF)
		indice1 := rand.Intn(lmotF)
		indice2 := rand.Intn(lmotF)
		indice3 := rand.Intn(lmotF)
		lettre1 := runesF[indice1]
		lettre2 := runesF[indice2]
		lettre3 := runesF[indice3]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmotF)
		}
		for indice1 == indice3 {
			indice3 = rand.Intn(lmotF)
		}
		for i, lettre := range runesF {
			if lettre == lettre1 || lettre == rune(lettre2) || lettre == rune(lettre3) {
				motstockF[i] = string(lettre)
			}
		}
		motstockF[indice1] = string(runesF[indice1])
		motstockF[indice2] = string(runesF[indice2])
		motstockF[indice3] = string(runesF[indice3])
	}
}

func HelpM() {
	if len(motM) <= 5 {
		rand.Seed(time.Now().UnixNano())
		lmotM := len(runesM)
		indice1 := rand.Intn(lmotM)
		lettre1 := runesM[indice1]
		motstockM[indice1] = string(runesM[indice1])
		for i, lettre := range runesM {
			if lettre == rune(lettre1) {
				motstockM[i] = string(lettre)
			}
		}
	} else if len(motM) <= 8 && len(motM) > 5 {
		rand.Seed(time.Now().UnixNano())
		lmotM := len(runesM)
		indice1 := rand.Intn(lmotM)
		indice2 := rand.Intn(lmotM)
		lettre1 := runesM[indice1]
		lettre2 := runesM[indice2]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmotM)
		}
		for i, lettre := range runesM {
			if lettre == rune(lettre1) || lettre == rune(lettre2) {
				motstockM[i] = string(lettre)
			}
		}
		motstockM[indice1] = string(runesM[indice1])
		motstockM[indice2] = string(runesM[indice2])
	} else if len(motF) < 8 {
		rand.Seed(time.Now().UnixNano())
		lmotM := len(runesM)
		indice1 := rand.Intn(lmotM)
		indice2 := rand.Intn(lmotM)
		indice3 := rand.Intn(lmotM)
		lettre1 := runesM[indice1]
		lettre2 := runesM[indice2]
		lettre3 := runesM[indice3]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmotM)
		}
		for indice1 == indice3 {
			indice3 = rand.Intn(lmotM)
		}
		for i, lettre := range runesM {
			if lettre == lettre1 || lettre == rune(lettre2) || lettre == rune(lettre3) {
				motstockM[i] = string(lettre)
			}
		}
		motstockM[indice1] = string(runesM[indice1])
		motstockM[indice2] = string(runesM[indice2])
		motstockM[indice3] = string(runesM[indice3])

	}

}

func HelpD() {
	if len(motD) <= 5 {
		rand.Seed(time.Now().UnixNano())
		lmotD := len(runesD)
		indice1 := rand.Intn(lmotD)
		lettre1 := runesD[indice1]
		motstockD[indice1] = string(runesD[indice1])
		for i, lettre := range runesD {
			if lettre == rune(lettre1) {
				motstockD[i] = string(lettre)
			}
		}
	} else if len(motD) <= 8 && len(motD) > 5 {
		rand.Seed(time.Now().UnixNano())
		lmotD := len(runesD)
		indice1 := rand.Intn(lmotD)
		indice2 := rand.Intn(lmotD)
		lettre1 := runesD[indice1]
		lettre2 := runesD[indice2]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmotD)
		}
		for i, lettre := range runesD {
			if lettre == rune(lettre1) || lettre == rune(lettre2) {
				motstockD[i] = string(lettre)
			}
		}
		motstockD[indice1] = string(runesD[indice1])
		motstockD[indice2] = string(runesD[indice2])
	} else if len(motF) < 8 {
		rand.Seed(time.Now().UnixNano())
		lmotD := len(runesD)
		indice1 := rand.Intn(lmotD)
		indice2 := rand.Intn(lmotD)
		indice3 := rand.Intn(lmotD)
		lettre1 := runesD[indice1]
		lettre2 := runesD[indice2]
		lettre3 := runesD[indice3]
		for indice1 == indice2 {
			indice2 = rand.Intn(lmotD)
		}
		for indice1 == indice3 {
			indice3 = rand.Intn(lmotD)
		}
		for i, lettre := range runesD {
			if lettre == lettre1 || lettre == rune(lettre2) || lettre == rune(lettre3) {
				motstockD[i] = string(lettre)
			}
		}
		motstockD[indice1] = string(runesD[indice1])
		motstockD[indice2] = string(runesD[indice2])
		motstockD[indice3] = string(runesD[indice3])
	}
}
