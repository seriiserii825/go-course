package main

import "fmt"

// 03. Server Storage Monitor
//
// Ты пишешь CLI-утилиту для проверки свободного
// пространства на Linux-сервере.
//
// Исходные данные:
// serverName := "Production Server"
// totalStorage := "500"
// usedStorage := "375.5"
// isOnline := "true"
//
// Задание:
// 1. Преобразуй totalStorage в int.
// 2. Преобразуй usedStorage в float64.
// 3. Преобразуй isOnline в bool.
// 4. Рассчитай свободное пространство.
// 5. Рассчитай процент использования диска.
// 6. Преобразуй процент использования в int.
// 7. Если использование превышает 70%,
//    выведи предупреждение.
// 8. Сформируй отчёт через fmt.Sprintf().
//
// Ожидаемый результат:
// === SERVER MONITOR ===
// Server: Production Server
// Online: true
//
// Total Storage: 500 GB
// Used Storage: 375.50 GB
// Free Storage: 124.50 GB
// Disk Usage: 75.10%
// Disk Usage (int): 75%
//
// WARNING: Disk usage is high!

func main() {
	_ = fmt.Sprintf
}
