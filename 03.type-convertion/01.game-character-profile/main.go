package main

import (
	"fmt"
	"strconv"
)

// 01. Game Character Profile
//
// Ты разрабатываешь CLI-приложение для RPG-игры.
// Данные персонажа приходят в виде строк, но для расчётов нужны числа.
//
// Исходные данные:
//
// Задание:
// 1. Преобразуй level из string в int.
// 2. Преобразуй health из string в float64.
// 3. Преобразуй isPremium из string в bool.
// 4. Вычисли максимальное здоровье: level * 10.
// 5. Преобразуй уровень обратно в строку через strconv.Itoa().
// 6. Создай строку профиля через fmt.Sprintf().
//
// Ожидаемый результат:
// === PLAYER PROFILE ===
// Name: Shadow
// Level: 15
// Health: 87.50
// Max Health: 150
// Premium: true

func main() {
	playerName := "Shadow"
	level := "15"
	health := "87.5"
	isPremium := "true"

	levelInt, err := strconv.Atoi(level)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	healthFloat, err := strconv.ParseFloat(health, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	isPremiumBool, err := strconv.ParseBool(isPremium)
	if err != nil {
		fmt.Println("Error", err)
	}
	maxHealth := levelInt * 10
	result := fmt.Sprintf("Name: %s\nLevel: %s\nHealth: %.2f\nMax Health: %d\nPremium: %t", playerName, level, healthFloat, maxHealth, isPremiumBool)
	fmt.Println(result)
}
