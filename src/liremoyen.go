package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var motM string

func MotAleatoireM(motsmoyens string) (string, error) {
	file, err := os.Open(motsmoyens)
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

func MotM() {
	motmoyen, err := MotAleatoireM("docs/motsmoyens.txt")
	if err != nil {
		log.Fatal(err)
	}
	motM = motmoyen
}
