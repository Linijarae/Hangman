package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var mot string

//Séléction d'un mot aléatoire facile
func MotAleatoire(motsfaciles string) (string, error) {
	file, err := os.Open(motsfaciles)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(mots) == 0 {
		return "", fmt.Errorf("le fichier est vide")
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(mots))
	return mots[randomIndex], nil
}

// Assignation à une variable de la liste facile
func MotF() {
	motfacile, err := MotAleatoire("docs/motsfaciles.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot = motfacile
}

// Assignation à une variable de la liste moyen
func MotM() {
	motmoyen, err := MotAleatoire("docs/motsmoyens.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot = motmoyen
}

// Assignation à une variable de la liste difficile
func MotD() {
	motdifficile, err := MotAleatoire("docs/motsdifficiles.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot = motdifficile
}

// Assignation à une variable de la liste motslongs
func MotLong() {
	motlong, err := MotAleatoire("docs/motslongs.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot = motlong
}

// Assignation à une variable de la liste medecine
func MotMed() {
	motmedecine, err := MotAleatoire("docs/motsmedecine.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot = motmedecine
}


