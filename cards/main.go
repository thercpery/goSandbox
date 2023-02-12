package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("myCards")

	cardsFromFile := newDeckFromFile("myCards")
	cardsFromFile.print()

	cards.shuffle()
	cards.print()

	// greeting := "I wish I was born in Norway"
	// fmt.Println([]byte(greeting)) // returns [73 32 119 105 115 104 32 73 32 119 97 115 32 98 111 114 110 32 105 110 32 78 111 114 119 97 121]

}
