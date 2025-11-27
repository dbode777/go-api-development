package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

const tableOdds = "3:2"

type Card struct {
	Value int // 1-13
	Suit  int // 0 = hearts, 1 = spades, 2 = clubs, 3 = diamonds
}

type Deck struct {
	Cards []Card
}

type Game struct {
	deck        Deck
	playerCards []Card
	dealerCards []Card
}

func (card Card) getCard() string {
	var suit string
	var value string

	switch card.Suit {
	case 0:
		suit = "❤️"
	case 1:
		suit = "♠️"
	case 2:
		suit = "♣️"
	case 3:
		suit = "♦️"
	}

	switch card.Value {
	case 1:
		value = "Ace"
	case 11:
		value = "Jack"
	case 12:
		value = "Queen"
	case 13:
		value = "King"
	default:
		value = strconv.Itoa(card.Value)
	}

	return value + suit
}

func hasBlackJack(cards []Card) bool {
	switch cards[0].Value {
	case 1:
		return cards[1].Value == 10 || cards[1].Value == 11 || cards[1].Value == 12 || cards[1].Value == 13
	case 10:
		return cards[1].Value == 1
	case 11:
		return cards[1].Value == 1
	case 12:
		return cards[1].Value == 1
	case 13:
		return cards[1].Value == 1
	default:
		return false
	}
}

func calculateValue(cards []Card) int {
	var value int
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value > cards[j].Value // sorting in descending order
	})
	for _, card := range cards {
		if card.Value > 10 {
			value += 10
		} else if card.Value == 1 && value <= 10 {
			value += 11
		} else {
			value += card.Value
		}
	}
	return value
}

func splitOption(cards []Card) bool {
	firstCardValue := cards[0].Value
	secondCardValue := cards[1].Value

	if firstCardValue > 10 {
		firstCardValue = 10
	}

	if secondCardValue > 10 {
		secondCardValue = 10
	}

	return firstCardValue == secondCardValue
}

func (deck *Deck) deal(num uint) []Card {
	var cards []Card
	for i := uint(0); i < num; i++ {
		cards = append(cards, deck.Cards[i])
	}
	deck.Cards = deck.Cards[num:]
	return cards
}

func (deck *Deck) create() {
	for suit := 0; suit < 4; suit++ {
		for value := 1; value < 14; value++ {
			deck.Cards = append(deck.Cards, Card{Value: value, Suit: suit})
		}
	}
}

func (deck *Deck) shuffle() {
	rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
}

func (game *Game) dealStartingCards() {
	game.playerCards = game.deck.deal(2)
	game.dealerCards = game.deck.deal(2)
}

func (game *Game) winConditions(cardValue int, bet float64) float64 {
	dealerValue := calculateValue(game.dealerCards)

	// Handles the case where no split occurs
	if cardValue == 0 {
		return 0
	}

	if cardValue > 21 {
		fmt.Println("You bust! You lose.")
		return -bet
	} else if dealerValue > 21 || dealerValue < cardValue {
		fmt.Println("The dealer busts or you are closer to 21! You win!")
		return bet
	} else if dealerValue > cardValue {
		fmt.Println("The dealer is closer to 21. You lose.")
		return -bet
	}

	fmt.Println("You and the dealer pushed. You get your bet back.")
	return 0
}

func (game *Game) play(bet float64) float64 {
	game.deck.create()
	game.deck.shuffle()
	game.dealStartingCards()
	fmt.Println("Dealer's first card: ", game.dealerCards[0].getCard())
	fmt.Println("Your cards: ", game.playerCards[0].getCard(), game.playerCards[1].getCard())

	// BlackJack conditions
	if hasBlackJack(game.playerCards) && hasBlackJack(game.dealerCards) {
		fmt.Println("You and the dealer both have BlackJack! This is a push and you get your bet back.")
		return 0
	} else if hasBlackJack(game.dealerCards) {
		fmt.Printf("The dealer has BlackJack with %s, %s! You lose.\n", game.dealerCards[0].getCard(), game.dealerCards[1].getCard())
		return -bet
	} else if hasBlackJack(game.playerCards) {
		fmt.Printf("You have BlackJack with %s, %s! You win on %s odds!\n", game.playerCards[0].getCard(), game.playerCards[1].getCard(), tableOdds)
		return bet * 1.5
	}

	fmt.Println("Would you like to double down? Press Y for yes or anything else to decline")
	input := enterInput()

	if input == "Y" || input == "y" {
		bet *= 2
	}

	splitValue := 0
	playerValue := 0
	if splitOption(game.playerCards) {
		fmt.Println("Would you like to split? Press Y for yes or anything else to decline")
		fmt.Println("Doubling down will split the cards and deal an additional card for each of the split cards, effectively giving you two hands to play with. An additional bet for the second handwill be required.")

		input := enterInput()
		if input == "Y" || input == "y" {
			playerCards := append([]Card{}, game.playerCards[0], game.deck.deal(1)[0])
			splitCards := append([]Card{}, game.playerCards[1], game.deck.deal(1)[0])

			// Will process the first one at the end of the play function
			fmt.Println("Playing the first hand")
			playerValue = game.playerTurn(playerCards)

			fmt.Println("Playing the split hand")
			splitValue = game.playerTurn(splitCards)
		} else {
			fmt.Println("You chose not to split.")
			playerValue = game.playerTurn(game.playerCards)
		}
	} else {
		playerValue = game.playerTurn(game.playerCards)
	}

	game.dealerTurn()
	return game.winConditions(splitValue, bet) + game.winConditions(playerValue, bet)
}

func (game *Game) playerTurn(cards []Card) int {
	fmt.Println("It's your turn!")
	var playerCardsString string
	var cardCount int
	playerCardsString += cards[0].getCard()
	playerCardsString += ", " + cards[1].getCard()
	game.playerCards = cards

	for {
		fmt.Println("Enter H to hit or S to stand.")
		cardCount = calculateValue(game.playerCards)
		fmt.Println("Your cards: ", playerCardsString, " with a value of", strconv.Itoa(cardCount))

		if cardCount == 21 {
			break
		}

		input := enterInput()
		if input == "H" || input == "h" {
			var dealtCard = game.deck.deal(1)[0]
			game.playerCards = append(game.playerCards, dealtCard)
			playerCardsString += ", " + dealtCard.getCard()
			cardCount = calculateValue(game.playerCards)
			if cardCount > 21 {
				fmt.Println("Your cards: ", playerCardsString, " with a value of", strconv.Itoa(cardCount))
				break
			}
		} else if input == "S" || input == "s" {
			break
		} else {
			fmt.Println("Invalid input. Please enter H to hit or S to stand.")
			continue
		}
	}

	return cardCount
}

func (game *Game) dealerTurn() {
	fmt.Println("Since you didn't bust, it's the dealer's turn!")
	var dealerCardsString string
	dealerCardsString += game.dealerCards[0].getCard()
	dealerCardsString += ", " + game.dealerCards[1].getCard()
	dealerCardCount := calculateValue(game.dealerCards)

	for dealerCardCount < 17 {
		var dealtCard = game.deck.deal(1)[0]
		game.dealerCards = append(game.dealerCards, dealtCard)
		dealerCardsString += ", " + dealtCard.getCard()
		dealerCardCount = calculateValue(game.dealerCards)
		fmt.Println("The dealer's cards after drawing a card: ", dealerCardsString, " with a value of", strconv.Itoa(dealerCardCount))
	}

	fmt.Println("The dealer's cards: ", dealerCardsString, "with a value of ", strconv.Itoa(dealerCardCount))
}

func enterInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // reads input after pressing enter on MacOS
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}

func main() {
	fmt.Printf("Welcom to BlackJack! The payout odds for blackjack are %s\n", tableOdds)
	balance := float64(100)

	for balance > 0 {
		fmt.Printf("Your balance is $%.2f\n", balance)
		fmt.Println("Enter your bet or enter Q to quit.")

		input := enterInput()
		if input == "Q" || input == "q" {
			fmt.Printf("Thanks for playing BlackJack! Your earnings for this session: $%.2f\n", balance-float64(100))
			break
		}
		bet, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid bet or enter Q to quit.")
			continue
		}

		if bet > balance || bet <= 5 {
			fmt.Println("Invalid bet. Please enter a valid bet or enter Q to quit. The minimum bet is $5 at this table.")
			continue
		}

		game := Game{}
		balance += game.play(bet)
	}
}
