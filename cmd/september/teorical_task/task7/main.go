// Что выведет программа и почему?
// Вы разрабатываете сервис для подсчета уникальных слов,
// но количество слов может быть очень большим,
// и вы решили ограничить максимальное число записей в мапе.
// Если в мапу добавляется больше слов, чем указано лимитом,
// она должна автоматически удалять самые старые записи.

package main

import "fmt"

type WordCounter struct {
	counts map[string]int
	limit  int

	order []string
}

func NewWordCounter(limit int) *WordCounter {
	return &WordCounter{
		counts: make(map[string]int),
		limit:  limit,
	}
}

func (wc *WordCounter) CountWord(word string) {
	_, ok := wc.counts[word]

	wc.counts[word]++

	if !ok {
		wc.order = append(wc.order, word)
	}

	if len(wc.counts) > wc.limit {
		// Логика удаления (здесь нужно реализовать)
		delete(wc.counts, wc.order[0])
		wc.order = append(wc.order[1:], word)

	}
}

func main() {
	wc := NewWordCounter(3)

	words := []string{"apple", "banana", "apple", "grape", "banana", "mango", "orange", "kiwi"}
	for _, word := range words {
		wc.CountWord(word)
		fmt.Println(wc.counts)
		fmt.Println(wc.order)
	}

	fmt.Println("Количество слов:", wc.counts)
}
