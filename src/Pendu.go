package Hangman

import (
	"fmt"
	"strings"
	"time"
)

// Si le joueur à choisi facile, facile passe en true
var facile = false

// Si le joueur à choisi moyen, moyen passe en true
var moyen = false

// Si le joueur à choisi difficile, difficile passe en true
var difficile = false

// Si le joueur à choisi les mots longs, motslongs passe en true
var motslongs = false

// Si le joueur à choisi medecine, medecine passe en true
var medecine = false
var runes []rune
var motstock []string
var passe = 0

func Pendu() {
	//Savoir quel niveau le joueur à séléctionné
	if facile {
		MotF()
	} else if moyen {
		MotM()
	} else if difficile {
		MotD()
	} else if motslongs {
		MotLong()
	} else if medecine {
		MotMed()
	}
	runes = []rune(mot)
	lmot := len(mot)
	motstock = make([]string, lmot)
	//Variable stock des input de lettres uniques
	var inputstock []string
	//Variable stock des "mots", c'est à dire les inputs d'une taille supérieure strictement à 1
	var inputstockmot []string
	//Contrôle si l'input correspond à une input précédemment rentré, pour éviter de rentrer dans une fonction que l'on ne souhaite pas qui serait vraie aussi.
	dejafait := false
	Underscore()
	Help()
	Prenom()
	//Boucle qui continue tant que motstock ne correspondra pas au mot
	for !strings.Contains(strings.Join(motstock, ""), mot) {
		dejafait = false
		ClearTerminal()
		//Affichage du schéma du pendu.
		Erreur()
		fmt.Println("║► \033[93mLe mot actuel :\033[0m", strings.Join(motstock, " "))
		fmt.Println("╠════════════════════════════════════════════════════════════■")
		fmt.Println("║► \033[93mVous avez utilisé les lettres :\033[0m", inputstock)
		fmt.Println("╠════════════════════════════════════════════════════════════■")
		fmt.Println("║► \033[93mVous avez tenté les mots :\033[0m", inputstockmot)
		fmt.Println("╠════════════════════════════════════════════════════════════■ ")
		fmt.Println("║")
		fmt.Printf("╚══█ \033[4m\033[1mVotre tentative : \033[0m")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		correcte := false
		//Input incorrecte changée en rouge
		redinput := "\033[91m" + input + "\033[0m"
		//Input correcte changée en verte
		greeninput := "\033[92m" + input + "\033[0m"
		for _, a := range inputstock {
			if a == redinput || a == greeninput {
				fmt.Println(" ")
				fmt.Println("\033[91m\033[1mVous avez déjà essayé cette lettre.\033[0m")
				fmt.Println(" ")
				dejafait = true
				time.Sleep(1500 * time.Millisecond)
				break
			}
		}
		// Si l'entrée utilisateur correspond au mot
		if mot == input {
			if facile {
				ReussiteF()
			} else if moyen {
				ReussiteM()
			} else if difficile {
				ReussiteD()
			} else if motslongs {
				ReussiteL()
			} else if medecine {
				ReussiteMed()
			}
			return
			// Le mot ( supérieur à 2 lettres ) entré par l'utilisateur ne correspond pas au mot recherché
		} else if mot != input && len(input) > 1 {
			fmt.Println(" ")
			fmt.Println("\033[91m\033[1mCe n'est pas le bon mot !\033[0m")
			fmt.Println(" ")
			time.Sleep(1500 * time.Millisecond)
			//Un mot faux c'est 2 erreurs en plus
			erreur += 2
			// Ajout du mot dans le stockmot en rouge car incorrect
			inputstockmot = append(inputstockmot, redinput)
		}
		//L'entrée utilisateur correspond à une lettre recherchée ET elle n'a pas encore été essayée
		for j, r := range runes {
			if string(r) == input && !dejafait {
				passe++
				motstock[j] = input
				correcte = true
				//Ajout de la lettre verte dans le stock de lettre, ajout d'une variable passe, pour n'afficher la lettre qu'une seule fois si elle est présente plusieurs fois.
				if passe == 1 {
					inputstock = append(inputstock, greeninput)
				}
			}
		}
		passe = 0
		// La lettre n'est pas bonne, et elle n'a pas déjà été tentée.
		if !correcte && len(input) == 1 && !dejafait {
			fmt.Println(" ")
			fmt.Println("\033[91m\033[1mLa lettre ne fait pas partie du mot !\033[0m")
			time.Sleep(1500 * time.Millisecond)
			//Ajout de la lettre rouge (car incorrecte) dans le stock de lettre
			inputstock = append(inputstock, redinput)
			//Incrémentation des erreurs
			erreur++
		}
	}
	//Toutes les lettres ont été trouvées, sortie de la boucle
	if facile {
		ReussiteF()
	} else if moyen {
		ReussiteM()
	} else if difficile {
		ReussiteD()
	} else if motslongs {
		ReussiteL()
	} else if medecine {
		ReussiteMed()
	}
}
