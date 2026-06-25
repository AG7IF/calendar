package ag7if

import (
	"time"

	cards "github.com/ag7if/playing-cards"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/ag7if/calendar/calc"
)

var deck cards.Deck = cards.Deck{
	{Suit: cards.Hearts, Rank: cards.Ace},
	{Suit: cards.Hearts, Rank: cards.Duce},
	{Suit: cards.Hearts, Rank: cards.Three},
	{Suit: cards.Hearts, Rank: cards.Four},
	{Suit: cards.Hearts, Rank: cards.Five},
	{Suit: cards.Hearts, Rank: cards.Six},
	{Suit: cards.Hearts, Rank: cards.Seven},
	{Suit: cards.Hearts, Rank: cards.Eight},
	{Suit: cards.Hearts, Rank: cards.Nine},
	{Suit: cards.Hearts, Rank: cards.Ten},
	{Suit: cards.Hearts, Rank: cards.Jack},
	{Suit: cards.Hearts, Rank: cards.Queen},
	{Suit: cards.Hearts, Rank: cards.King},

	{Suit: cards.Diamonds, Rank: cards.Ace},
	{Suit: cards.Diamonds, Rank: cards.Duce},
	{Suit: cards.Diamonds, Rank: cards.Three},
	{Suit: cards.Diamonds, Rank: cards.Four},
	{Suit: cards.Diamonds, Rank: cards.Five},
	{Suit: cards.Diamonds, Rank: cards.Six},
	{Suit: cards.Diamonds, Rank: cards.Seven},
	{Suit: cards.Diamonds, Rank: cards.Eight},
	{Suit: cards.Diamonds, Rank: cards.Nine},
	{Suit: cards.Diamonds, Rank: cards.Ten},
	{Suit: cards.Diamonds, Rank: cards.Jack},
	{Suit: cards.Diamonds, Rank: cards.Queen},
	{Suit: cards.Diamonds, Rank: cards.King},

	{Suit: cards.Clubs, Rank: cards.Ace},
	{Suit: cards.Clubs, Rank: cards.Duce},
	{Suit: cards.Clubs, Rank: cards.Three},
	{Suit: cards.Clubs, Rank: cards.Four},
	{Suit: cards.Clubs, Rank: cards.Five},
	{Suit: cards.Clubs, Rank: cards.Six},
	{Suit: cards.Clubs, Rank: cards.Seven},
	{Suit: cards.Clubs, Rank: cards.Eight},
	{Suit: cards.Clubs, Rank: cards.Nine},
	{Suit: cards.Clubs, Rank: cards.Ten},
	{Suit: cards.Clubs, Rank: cards.Jack},
	{Suit: cards.Clubs, Rank: cards.Queen},
	{Suit: cards.Clubs, Rank: cards.King},

	{Suit: cards.Spades, Rank: cards.Ace},
	{Suit: cards.Spades, Rank: cards.Duce},
	{Suit: cards.Spades, Rank: cards.Three},
	{Suit: cards.Spades, Rank: cards.Four},
	{Suit: cards.Spades, Rank: cards.Five},
	{Suit: cards.Spades, Rank: cards.Six},
	{Suit: cards.Spades, Rank: cards.Seven},
	{Suit: cards.Spades, Rank: cards.Eight},
	{Suit: cards.Spades, Rank: cards.Nine},
	{Suit: cards.Spades, Rank: cards.Ten},
	{Suit: cards.Spades, Rank: cards.Jack},
	{Suit: cards.Spades, Rank: cards.Queen},
	{Suit: cards.Spades, Rank: cards.King},

	{Suit: cards.Black, Rank: cards.Joker},
}

func ComputeWeekPlayingCard(isoweek int) (cards.Card, error) {
	if isoweek < 1 || isoweek > 53 {
		return cards.Card{}, errors.Errorf("invaild week number: %d", isoweek)
	}

	return deck[isoweek-1], nil
}

func WeekdayLetter(wd time.Weekday) string {
	switch wd {
	case time.Monday:
		return "M"
	case time.Tuesday:
		return "T"
	case time.Wednesday:
		return "W"
	case time.Thursday:
		return "H"
	case time.Friday:
		return "F"
	case time.Saturday:
		return "S"
	case time.Sunday:
		return "U"
	default:
		panic(errors.Errorf("invalid weekday value: %d", wd))
	}
}

func ComputeWeek1StartDate(year int) date.Date {
	start := calc.ComputeNearestMonday(date.New(year, time.January, 1))

	tyr, wk := start.ISOWeek()
	if tyr < year {
		return start.Add(7)
	}

	if wk > 1 {
		return start.Add(-7)
	}

	return start
}
