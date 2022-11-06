package service

import (
	"io/ioutil"
	"strings"
)

func ReadFile(fileName string) ([]string, error) {
	fileRead, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	tempSlice := string(fileRead)
	deduplicate := unique(strings.Split(tempSlice[:len(tempSlice)-1], ","))
	return deduplicate, nil

}

func unique(cards []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range cards {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
