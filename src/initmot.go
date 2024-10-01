package Hangman

// Initialisation du mot facile avec des underscores bleus
func UnderscoreF() {
	for i := range motstockF {
		motstockF[i] = "\033[94m_\033[0m"
	}
}
// Initialisation du mot moyen avec des underscores bleus
func UnderscoreM() {
	for i := range motstockM {
		motstockM[i] = "\033[94m_\033[0m"
	}
}
// Initialisation du mot difficile avec des underscores bleus
func UnderscoreD() {
	for i := range motstockD {
		motstockD[i] = "\033[94m_\033[0m"
	}
}
