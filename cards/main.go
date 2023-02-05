package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("myCards")

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting)) // returns [72 105 32 116 104 101 114 101 33]
}
