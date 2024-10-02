package Hangman

// Initialisation du mot facile avec des underscores bleus
func Underscore() {
	for i := range motstock {
		motstock[i] = "\033[94m_\033[0m"
	}
}
