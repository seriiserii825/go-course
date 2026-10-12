package main

import (
	"fmt"
	"strings"
)

// Ты разрабатываешь систему регистрации пользователей.
// Исходные данные:
// username := "  Sergey_Dev  "
// Задание:
// Удали пробелы в начале и конце.
// Преобразуй username в нижний регистр.
// Проверь, что длина имени от 5 до 20 символов.
// Проверь, содержит ли имя "admin".
// Проверь, начинается ли имя с "ser".
// Выведи результат через fmt.Sprintf().
// Ожидаемый результат:
// === USERNAME VALIDATOR ===
// Username: sergey_dev
// Length: 10
// Valid Length: true
// Contains Admin: false
// Starts With Ser: true
func main() {
	username := "  sergey_administrator_for_blueline  "
	usernameWithoutSpaces := strings.TrimSpace(username)
	wordLength := len(usernameWithoutSpaces)
	const tresholdFrom = 5
	const tresholdTo = 20
	validLength := false
	if wordLength >= tresholdFrom && wordLength <= tresholdTo {
		validLength = true
	}
	containsAdmin := strings.Contains(usernameWithoutSpaces, "admin")
	startsWithWord := strings.HasPrefix(usernameWithoutSpaces, "ser")
	result := fmt.Sprintf(`=== USERNAME VALIDATOR ===
	Username: %s
	Length: %d
	Valid Length: %t
	Contains Admin: %t
	Starts With Ser: %t
	`, usernameWithoutSpaces, wordLength, validLength, containsAdmin, startsWithWord)
	fmt.Println(result)
}
