package Hangman

func UnderscoreF() {
	for i := range motstockF {
		motstockF[i] = "\033[94m_\033[0m"
	}
}
func UnderscoreM() {
	for i := range motstockM {
		motstockM[i] = "\033[94m_\033[0m"
	}
}

func UnderscoreD() {
	for i := range motstockD {
		motstockD[i] = "\033[94m_\033[0m"
	}
}
