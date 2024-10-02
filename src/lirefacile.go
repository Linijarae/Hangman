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
func MotAleatoireF(motsfaciles string) (string, error) {
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

// Assignation à une variable du mot facile
func MotF() {
	motfacile, err := MotAleatoireF("docs/motsfaciles.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot = motfacile
}
