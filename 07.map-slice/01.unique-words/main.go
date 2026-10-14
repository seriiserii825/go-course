package main

import "fmt"

// # Проект 1: Уникальные слова
// ## Задача
// Написать функцию, которая принимает `[]string` (список слов) и возвращает новый `[]string` только с уникальными словами, сохраняя порядок первого появления.
// Используй `map[string]bool` (или `map[string]struct{}`) для отслеживания уже встреченных слов и `slice` для накопления результата.
// Дополнительно: написать вторую функцию, которая считает, сколько раз встретилось каждое слово, и возвращает `map[string]int`.
// ## Результат
// - Функция `UniqueWords(words []string) []string` — убирает повторы, порядок первого появления сохранён.
// - Функция `WordCount(words []string) map[string]int` — ключ — слово, значение — количество повторений.
// - Написать `main()`, который демонстрирует работу обеих функций на любом наборе слов (минимум 10 слов, с повторами).
// - Вывести в консоль результат обеих функций через `fmt.Println`.
func main() {
	words := []string{
		"Alex", "Serii", "Nixon", "Serii", "Anna",
		"Max", "Alex", "Olga", "Nixon", "Ivan",
		"Anna", "Petr", "Max", "Olga", "Ivan",
		"Serii", "Petr", "Alex", "Anna", "Nixon",
	}
	UniqueWords(words)
}

func UniqueWords(words []string) []string {
	fmt.Println(words)
	uniquesWords := []string{}
	wordsMap := map[string]bool{}

	for _, word := range words {
		_, exists := wordsMap[word]
		if !exists {
			uniquesWords = append(uniquesWords, word)
		}
		wordsMap[word] = true
	}
	fmt.Println(wordsMap)
	fmt.Println(uniquesWords)
	return uniquesWords
}
