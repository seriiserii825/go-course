package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Ты пишешь CLI-утилиту для анализа email-адресов.
// Исходные данные:
// Задание:
// Удали лишние пробелы.
// Преобразуй email в нижний регистр.
// Проверь, содержит ли email символ @.
// Раздели email на username и domain через strings.Split().
// Проверь, равен ли домен "gmail.com".
// Замени точки в username на _.
// Сформируй новый email.
// Выведи отчёт.
// Ожидаемый результат:
// === EMAIL ANALYZER ===
// Original: Sergey.Dev@GMAIL.COM
// Normalized: sergey.dev@gmail.com
// Username: sergey.dev
// Domain: gmail.com
// Is Gmail: true
// New Username: sergey_dev
// New Email: sergey_dev@gmail.com
func main() {
	email := "  Sergey.Dev@GMAIL.COM  "
	emailWithoutSpaces := strings.TrimSpace(email)
	emailLower := strings.ToLower(emailWithoutSpaces)
	isEmail := strings.Contains(emailLower, "@")
	if !isEmail {
		color.Blue("String is not email")
		return
	}
	emailOfStrings := strings.Split(emailLower, "@")
	userName := emailOfStrings[0]
	domain := emailOfStrings[1]
	isGmail := domain == "gmail.com"
	newUserName := strings.ReplaceAll(userName, ".", "_")
	new_email := fmt.Sprintf("%s@%s", newUserName, domain)

	result := fmt.Sprintf(`=== Email Analyzer ===
  Original: %s
  Normalized: %s
  Username: %s
  Domain: %s
  Is Gmail: %t
  New Username: %s
  New Email: %s
  `,
		email,
		emailLower,
		userName,
		domain,
		isGmail,
		newUserName,
		new_email,
	)
	fmt.Println(result)
}
