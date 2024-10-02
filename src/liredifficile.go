package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)


//Séléction d'un mot aléatoire difficile
func MotAleatoireD(motsdifficile string) (string, error) {
	file, err := os.Open(motsdifficile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(words) == 0 {
		return "", fmt.Errorf("le fichier est vide")
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	return words[randomIndex], nil
}
// Assignation à une variable du mot difficile
func MotD() {
	motdifficile, err := MotAleatoireD("docs/motsdifficiles.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot = motdifficile
}
