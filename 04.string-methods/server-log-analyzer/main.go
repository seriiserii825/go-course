package main

import "fmt"

// Ты разрабатываешь CLI-инструмент для анализа логов Linux-сервера.
// Исходные данные:
// log := "  ERROR: Database connection failed; WARNING: Disk usage 85%; INFO: Server restarted; ERROR: Timeout  "
// Задание:
// Удали пробелы по краям строки.
// Раздели лог на отдельные сообщения по ;.
// Удали пробелы у каждого сообщения.
// Посчитай количество сообщений.
// Посчитай количество ERROR, WARNING, INFO.
// Проверь, есть ли ошибка подключения к базе данных.
// Замени "ERROR" на "CRITICAL".
// Объедини все сообщения обратно через strings.Join() с разделителем " | ".
// Выведи итоговый отчёт через fmt.Sprintf().
// Ожидаемый результат:
// === SERVER LOG ANALYZER ===
// Total Messages: 4
// Errors: 2
// Warnings: 1
// Info: 1
// Database Error: true
// Formatted Logs:
// CRITICAL: Database connection failed
// WARNING: Disk usage 85%
// INFO: Server restarted
// CRITICAL: Timeout
// Single Line:
// CRITICAL: Database connection failed | WARNING: Disk usage 85% | INFO: Server restarted | CRITICAL: Timeout
func main() {
}
