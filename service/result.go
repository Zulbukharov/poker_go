package service

import (
	"github.com/manarakozhamuratova/pokerGo/combination"
)

func GetAnswerPokerCombination(filepath string, answerCh chan<- Answer, i int) {
	cardsStrSlice, err := ReadFile(filepath)
	if err != nil {
		// errCh <- err
		return
	}

	allComb := combination.CombinationCards(cardsStrSlice)

	var cardCombinationAnswer string

	for _, comb := range allComb {
		cardComb, err := combination.New(comb)
		if err != nil {
			// errCh <- err
			return
		}
		if err := cardComb.IsPokerCombination(); err != nil {
			continue
		}
		cardCombinationAnswer += cardComb.String()
		//fmt.Println(cardCombinationAnswer)
	}

	answer := Answer{
		filePath: filepath,
		answer:   cardCombinationAnswer,
		index:    i,
	}
	answerCh <- answer
}
