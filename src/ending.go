package Hangman

import (
	"fmt"
	"time"
)

var killtotal = 0

// Fin du jeu en cas d'arrivée à 150 points
func Ending() {
	ClearTerminal()
	killtotal = killF + killM + killD + killL + killmed
	fmt.Println(" ")
	fmt.Println("\033[1mVous avez atteint les \033[92m", limitescore, "\033[0m\033[1m points !!!!")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\033[92m\033[1mFélicitations !!!\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\033[93m\033[1mL'homme rentre dans la pièce. Il éteint l'écran en face de vous.\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\033[95m\033[1m??????\033[0m - Eh bien, vous avez tout de même fait mourir\033[91m", killtotal, "\033[0mpersonnes ...")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\033[95m\033[1m??????\033[0m - Ce n'est pas si mauvais, mes patients précédents étaient bien moins bon.")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\033[95m\033[1m??????\033[0m - Et bien vous pouvez partir.")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\033[1m\033[93mL'homme ouvre une porte derrière vous. Il vous détache et quitte la pièce.")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Vous vous levez et sortez par la porte qu'il a ouverte, vous êtes ébloui par la lumière.\033[0m")
	fmt.Println(" ")
	time.Sleep(1000 * time.Millisecond)
	ClearTerminal()
	Flashbang()
	fmt.Println(" ")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println(" ")
	ClearTerminal()
	fmt.Println("\033[1m\033[96m                                MERCI D'AVOIR JOUER                                \033[0m")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(" ")
	Quitter()
}
