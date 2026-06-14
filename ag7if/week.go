package ag7if

import (
	cards "github.com/ag7if/playing-cards"
	"github.com/pkg/errors"
)

var deck cards.Deck = cards.Deck{
	{cards.Hearts, cards.Ace},
	{cards.Hearts, cards.Duce},
	{cards.Hearts, cards.Three},
	{cards.Hearts, cards.Four},
	{cards.Hearts, cards.Five},
	{cards.Hearts, cards.Six},
	{cards.Hearts, cards.Seven},
	{cards.Hearts, cards.Eight},
	{cards.Hearts, cards.Nine},
	{cards.Hearts, cards.Ten},
	{cards.Hearts, cards.Jack},
	{cards.Hearts, cards.Queen},
	{cards.Hearts, cards.King},

	{cards.Diamonds, cards.Ace},
	{cards.Diamonds, cards.Duce},
	{cards.Diamonds, cards.Three},
	{cards.Diamonds, cards.Four},
	{cards.Diamonds, cards.Five},
	{cards.Diamonds, cards.Six},
	{cards.Diamonds, cards.Seven},
	{cards.Diamonds, cards.Eight},
	{cards.Diamonds, cards.Nine},
	{cards.Diamonds, cards.Ten},
	{cards.Diamonds, cards.Jack},
	{cards.Diamonds, cards.Queen},
	{cards.Diamonds, cards.King},

	{cards.Clubs, cards.Ace},
	{cards.Clubs, cards.Duce},
	{cards.Clubs, cards.Three},
	{cards.Clubs, cards.Four},
	{cards.Clubs, cards.Five},
	{cards.Clubs, cards.Six},
	{cards.Clubs, cards.Seven},
	{cards.Clubs, cards.Eight},
	{cards.Clubs, cards.Nine},
	{cards.Clubs, cards.Ten},
	{cards.Clubs, cards.Jack},
	{cards.Clubs, cards.Queen},
	{cards.Clubs, cards.King},

	{cards.Spades, cards.Ace},
	{cards.Spades, cards.Duce},
	{cards.Spades, cards.Three},
	{cards.Spades, cards.Four},
	{cards.Spades, cards.Five},
	{cards.Spades, cards.Six},
	{cards.Spades, cards.Seven},
	{cards.Spades, cards.Eight},
	{cards.Spades, cards.Nine},
	{cards.Spades, cards.Ten},
	{cards.Spades, cards.Jack},
	{cards.Spades, cards.Queen},
	{cards.Spades, cards.King},

	{cards.Black, cards.Joker},
}

func ComputeWeekPlayingCard(isoweek int) (cards.Card, error) {
	if isoweek < 1 || isoweek > 53 {
		return cards.Card{}, errors.Errorf("invaild week number: %d", isoweek)
	}

	return deck[isoweek-1], nil
}
