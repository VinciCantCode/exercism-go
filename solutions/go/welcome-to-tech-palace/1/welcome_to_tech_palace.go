package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	firstLine := strings.Repeat("*", numStarsPerLine)
	secoundLine := welcomeMsg
	thirdLine := strings.Repeat("*",numStarsPerLine)
	return firstLine + "\n" +secoundLine + "\n" + thirdLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg,"*",""))
}
