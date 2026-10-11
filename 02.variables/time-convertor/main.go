package main

import "fmt"

// Задание 4. Конвертер времени
// Уровень 4/5
// Программа получает количество секунд и преобразует их в часы, минуты и секунды.
// totalSeconds := 93784
// Ожидаемый результат:
// === Time Converter ===
// Input: 93784 seconds
// Days: 1
// Hours: 2
// Minutes: 3
// Seconds: 4
// Подсказка: используй деление / и остаток от деления %.
// Например:
// minutes := 125 / 60
// seconds := 125 % 60
// fmt.Println(minutes) // 2
// fmt.Println(seconds) // 5
// Здесь нужно самому придумать последовательность вычислений.
func main() {
	totalSeconds := 350730
	const hoursInDay = 24
	const minutesInHour = 60
	const secondsInMinute = 60
	const daySeconds = hoursInDay * minutesInHour * secondsInMinute
	days := totalSeconds / (daySeconds)
	daysSecondsRemaining := totalSeconds % daySeconds
	hours := daysSecondsRemaining / (minutesInHour * secondsInMinute)
	hoursSecondsRemaining := daysSecondsRemaining % (minutesInHour * secondsInMinute)
	minutes := hoursSecondsRemaining / secondsInMinute
	seconds := hoursSecondsRemaining % secondsInMinute

	fmt.Println("=== Time Converter ===")
	fmt.Printf("Input: %d seconds\n", totalSeconds)
	fmt.Printf("Days: %d\n", days)
	fmt.Printf("Hours: %d\n", hours)
	fmt.Printf("Minutes: %d\n", minutes)
	fmt.Printf("Seconds: %d\n", seconds)
}
