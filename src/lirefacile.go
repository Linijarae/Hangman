package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var motF string

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

func MotF() {
	motfacile, err := MotAleatoireF("motsfaciles.txt")
	if err != nil {
		log.Fatal(err)
	}
	motF = motfacile
}
