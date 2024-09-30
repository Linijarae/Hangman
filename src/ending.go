package Hangman

import (
	"fmt"
	"time"
)

var killtotal = 0

func Ending() {
		killtotal = killF + killM + killD
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("\033[1mVous avez atteint les \033[92m150\033[0m\033[1m points !!!!")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Félicitations !!!")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("L'homme rentre dans la pièce. Il éteint l'écran en face de vous.\033[0m")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("?????? - Et bien, vous avez tout de même fait mourir", killtotal, "personnes ...")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("?????? - Ce n'est pas si mauvais, mes patients précédents étaient bien moins bon.")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("?????? - Et bien vous pouvez partir.")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("\033[1mL'homme ouvre une porte derrière vous. Il vous détache et quitte la pièce.")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Vous vous levez et sortez par la porte qu'il a ouverte, vous êtes ébloui par la lumière.\033[0m")
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		Flashbang()
		fmt.Println(" ")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("\033[1mMERCI D'AVOIR JOUER\033[0m")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(" ")
		Quitter()
}
