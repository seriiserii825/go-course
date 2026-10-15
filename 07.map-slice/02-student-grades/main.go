package main

import "fmt"

// # Проект 2: Журнал оценок студентов
// ## Задача
// Смоделировать журнал оценок, используя `map[string][]int`, где ключ — имя студента, а значение — slice его оценок.
// Реализовать функции:
// - добавление оценки студенту (если студента нет в map — создать его);
// - вычисление среднего балла одного студента;
// - определение студента с самым высоким средним баллом среди всех;
// - получение списка имён всех студентов, отсортированного по алфавиту (использовать `sort` или пакет `slices`).
// ## Результат
// - Функция `AddGrade(grades map[string][]int, name string, grade int)` — добавляет оценку.
// - Функция `AverageGrade(grades map[string][]int, name string) float64` — средний балл студента.
// - Функция `TopStudent(grades map[string][]int) string` — имя студента с наибольшим средним баллом.
// - Функция `SortedNames(grades map[string][]int) []string` — имена студентов по алфавиту.
// - `main()` с минимум 3 студентами и несколькими оценками у каждого, демонстрирующий работу всех функций через `fmt.Println`.
func main() {
	grades := map[string][]int{
		"John": {4, 8, 4},
		"Anna": {3, 2, 9},
	}
	fmt.Println(grades)
	AddGrade(grades, "Fill", 9)
	fmt.Println(grades)
	userName := "John"
	johnGrade := AverageGrade(grades, userName)
	fmt.Printf("Average %s: %.2f\n", userName, johnGrade)
}

func AddGrade(grades map[string][]int, name string, grade int) {
	_, exists := grades[name]
	if !exists {
		grades[name] = []int{grade}
	} else {
		grades[name] = append(grades[name], grade)
	}
}

func AverageGrade(grades map[string][]int, name string) float64 {
	userGrades := grades[name]
	userGradesLen := len(userGrades)
	if userGradesLen == 0 {
		return 0
	}
	sum := 0
	for _, g := range userGrades {
		sum += g
	}
	return float64(sum) / float64(userGradesLen)
}
