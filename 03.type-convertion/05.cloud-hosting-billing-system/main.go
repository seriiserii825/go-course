package main

import "fmt"

// 05. Cloud Hosting Billing System
//
// Ты разрабатываешь систему расчёта стоимости
// облачного сервера.
//
// Клиент оплачивает CPU, RAM, дисковое пространство
// и дополнительные услуги.
//
// Исходные данные:
// serverName := "cloud-prod-01"
//
// cpuCores := "8"
// cpuPrice := "12.50"
//
// ramGB := "32"
// ramPrice := "3.75"
//
// storageGB := "250"
// storagePrice := "0.15"
//
// backupEnabled := "true"
// backupPrice := "19.99"
//
// discount := "10"
// billingMonths := 3
//
// Задание:
// 1. Преобразуй все числовые строки
//    в подходящие типы.
// 2. Преобразуй backupEnabled в bool.
// 3. Рассчитай месячную стоимость CPU.
// 4. Рассчитай месячную стоимость RAM.
// 5. Рассчитай месячную стоимость Storage.
// 6. Если backupEnabled == true,
//    добавь стоимость резервного копирования.
// 7. Рассчитай общую месячную стоимость.
// 8. Примени скидку 10%.
// 9. Рассчитай итоговую стоимость за 3 месяца.
// 10. Преобразуй итоговую стоимость в строку
//     через strconv.FormatFloat().
// 11. Сформируй отчёт через fmt.Sprintf().
//
// Ожидаемый результат:
// === CLOUD HOSTING BILL ===
// Server: cloud-prod-01
//
// CPU: 8 cores x $12.50 = $100.00
// RAM: 32 GB x $3.75 = $120.00
// Storage: 250 GB x $0.15 = $37.50
// Backup: $19.99
//
// Monthly Subtotal: $277.49
// Discount (10%): $27.75
// Monthly Total: $249.74
//
// Billing Period: 3 months
// Final Total: $749.22

func main() {
	_ = fmt.Sprintf
}
