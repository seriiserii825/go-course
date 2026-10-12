package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

// 03. Server Storage Monitor
//
// Ты пишешь CLI-утилиту для проверки свободного
// пространства на Linux-сервере.
//
// Исходные данные:
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
	const serverName = "Production Server"
	totalStorage := "400"
	usedStorage := "375.5"
	isOnline := "true"

	totalStorageInt, err := strconv.Atoi(totalStorage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	usedStorageFloat, err := strconv.ParseFloat(usedStorage, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	isOnlineBool, err := strconv.ParseBool(isOnline)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	freeSpace := float64(totalStorageInt) - usedStorageFloat
	usagePercent := (usedStorageFloat * 100) / float64(totalStorageInt)
	usagePercentInt := int(usagePercent)
	fmt.Println("=== SERVER MONITOR ===")
	result := fmt.Sprintf(`Server: %s 
	Online: %t 

	Total Storage: %d
	GB Used Storage: %.2f
	GB Free Storage: %.2f
	Disk Usage: %.2f%% 
	Disk Usage (int): %d%%
	`,
		serverName,
		isOnlineBool,
		totalStorageInt,
		usedStorageFloat,
		freeSpace,
		usagePercent,
		usagePercentInt,
	)
	fmt.Println(result)

	const warningMessage = "WARNING: Disk usage is high!"
	warningThreshold := 70.0
	if usagePercent > warningThreshold {
		color.Red(warningMessage)
	}
}
