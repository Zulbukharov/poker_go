package combination

import (
	"fmt"

	"github.com/manarakozhamuratova/pokerGo/card"
)

type Cards []card.Card

func (cards Cards) String() string {
	var result string

	for i := range cards {
		if i == len(cards)-1 {
			result += cards[i].String()
		} else {
			result += fmt.Sprintf("%s,", cards[i].String())
		}
	}
	return result
}

func (cards Cards) Len() int {
	return len(cards)
}

func (cards Cards) Less(i, j int) bool {
	return numFace(cards[i].Face) > numFace(cards[j].Face)
}

func (cards Cards) Swap(i, j int) {
	cards[i], cards[j] = cards[j], cards[i]
}
