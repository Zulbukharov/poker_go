package service

type Answer struct {
	filePath string
	answer   string
	index    int
}

func (a Answer) GetAnswer() string {
	return a.answer
}

func (a Answer) Index() int {
	return a.index
}
