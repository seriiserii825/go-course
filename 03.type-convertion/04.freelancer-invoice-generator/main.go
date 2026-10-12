package main

import "fmt"

// 04. Freelancer Invoice Generator
//
// Ты создаёшь CLI-программу для генерации
// счетов клиентам.
//
// Исходные данные:
// clientName := "Tech Solutions"
// hourlyRate := "45.50"
// hoursWorked := "38"
// taxRate := "20"
// invoicePaid := "false"
// invoiceNumber := 1042
//
// Задание:
// 1. Преобразуй hourlyRate в float64.
// 2. Преобразуй hoursWorked в int.
// 3. Преобразуй taxRate в int, затем в float64.
// 4. Преобразуй invoicePaid в bool.
// 5. Рассчитай стоимость работы без налога.
// 6. Рассчитай сумму налога.
// 7. Рассчитай итоговую стоимость.
// 8. Преобразуй invoiceNumber в строку
//    через strconv.Itoa().
// 9. Преобразуй итоговую сумму в строку
//    через strconv.FormatFloat().
// 10. Сформируй счёт через fmt.Sprintf().
//
// Дополнительно:
// Если invoicePaid == true, выведи Status: PAID.
// Иначе выведи Status: PENDING.
//
// Ожидаемый результат:
// === INVOICE #1042 ===
// Client: Tech Solutions
//
// Hourly Rate: $45.50
// Hours Worked: 38
// Subtotal: $1729.00
// Tax (20%): $345.80
//
// Total: $2074.80
// Paid: false
// Status: PENDING

func main() {
	_ = fmt.Sprintf
}
