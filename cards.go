package main

func main() {
	//cards := newDeck()

	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	//cards:= newDeckFromFile("my cards") //For newDeckofFlie
	//cards.print()
	cards := newDeck()//For Shuffle
	cards.shuffle()
	cards.print()
}
