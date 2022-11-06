package combination

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/manarakozhamuratova/pokerGo/card"
	combo "github.com/mxschmitt/golang-combinations"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

func CombinationCards(cards []string) [][]string {
	return combo.Combinations(cards, CombinationSize)
}

const (
	CombinationSize          = 5
	CombinationPairName      = "Pair"
	CombinationTwoPairs      = "Two Pairs"
	CombinationThreeOfAKind  = "Three Of A Kind"
	CombinationStraight      = "Straight"
	CombinationFlush         = "Flush"
	CombinationFullHouse     = "Full House"
	CombinationFourOfAKind   = "Four Of A Kind"
	CombinationStraightFlush = "Straight Flush"
	NumericJack              = 11
	NumericQueen             = 12
	NumericKing              = 13
	NumericAce               = 14
)

var (
	ErrMissingName  = errors.New("missing combination name")
	ErrMissingCards = errors.New("missing cards")
)

type Combination struct {
	name  string
	cards Cards
}

func New(cards []string) (*Combination, error) {
	if cards == nil {
		return nil, ErrMissingCards
	}
	comb := &Combination{
		name:  "",
		cards: make([]card.Card, 0),
	}
	for _, w := range cards {
		tmpCard, err := card.NewWithUnicode(string([]rune(w)[0]), string([]rune(w)[1:]))
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert to card")
		}
		comb.cards = append(comb.cards, *tmpCard)
	}
	return comb, nil
}

func (c Combination) Name() string {
	return c.name
}

func (c Combination) String() string {
	return fmt.Sprintf("%s | %s\n", c.cards, c.name)
}

func countFaces(cards []card.Card) map[string]int {
	faceCount := map[string]int{}
	for _, card := range cards {
		if count, ok := faceCount[card.Face]; ok {
			faceCount[card.Face] = count + 1
		} else {
			faceCount[card.Face] = 1
		}
	}
	return faceCount
}

func sortValues(mapComb map[string]int) []int {
	values := lo.Values[string, int](mapComb)
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	return values
}

func isCombOfPair(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 2
}

func isCombOfTwoPair(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 2 && sortValues(countFaces(cards))[1] == 2
}

func isCombOfThreeOfAKind(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 3
}

func numFace(cards string) int {
	switch cards {
	case card.FaceJack:
		return NumericJack
	case card.FaceQueen:
		return NumericQueen
	case card.FaceKing:
		return NumericKing
	case card.FaceAce:
		return NumericAce
	default:
		num, _ := strconv.Atoi(cards)
		return num
	}
}

func isCombOfStraight(cards Cards) bool {
	var sliceFace []int

	for _, value := range cards {
		sliceFace = append(sliceFace, numFace(value.Face))
	}

	// sort.Sort(cards)
	sort.Slice(sliceFace, func(i, j int) bool { return sliceFace[i] > sliceFace[j] })
	// log.Println("native", cards, "string", sliceFace)

	if slices.Contains(sliceFace, NumericAce) && slices.Contains(sliceFace, 2) {
		return sliceFace[0] == NumericAce && sliceFace[1] == 5 && sliceFace[2] == 4 && sliceFace[3] == 3 && sliceFace[4] == 2
	} else {
		for i := 0; i < len(sliceFace); i++ {
			if i != len(sliceFace)-1 && sliceFace[i]-sliceFace[i+1] != 1 {
				return false
			}
		}
	}
	return true
}

func isCombOfFlush(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 5
}

func isCombFullHouse(cards []card.Card) bool {
	return isCombOfThreeOfAKind(cards) && isCombOfPair(cards)
}

func isCombFourOfAKind(cards []card.Card) bool {
	return sortValues(countFaces(cards))[0] == 4
}

func isCombOfStraightFlush(cards []card.Card) bool {
	return isCombOfStraight(cards) && isCombOfFlush(cards)
}

func (c *Combination) IsPokerCombination() error {
	if len(c.cards) != 5 {
		return errors.New("Not correct")
	}
	switch {
	case isCombOfStraightFlush(c.cards):
		c.name = CombinationStraightFlush
	case isCombFourOfAKind(c.cards):
		c.name = CombinationFourOfAKind
	case isCombFullHouse(c.cards):
		c.name = CombinationFullHouse
	case isCombOfFlush(c.cards):
		c.name = CombinationFlush
	case isCombOfStraight(c.cards):
		c.name = CombinationStraight
	case isCombOfThreeOfAKind(c.cards):
		c.name = CombinationThreeOfAKind
	case isCombOfTwoPair(c.cards):
		c.name = CombinationTwoPairs
	case isCombOfPair(c.cards):
		c.name = CombinationPairName
	default:
		return errors.New("Not combination")
	}
	return nil
}
