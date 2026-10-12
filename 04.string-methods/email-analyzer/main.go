package main

import "fmt"

// Ты пишешь CLI-утилиту для анализа email-адресов.
// Исходные данные:
// email := "  Sergey.Dev@GMAIL.COM  "
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
}
