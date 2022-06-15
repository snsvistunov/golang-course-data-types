// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Під опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані // відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

package main

import "fmt"

func main() {

	var (
		priceOfOneApple float64 = 5.99
		priceOfOnePear  float64 = 7.
		totalAmount             = 23.
	)

	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?

	numberOfApplesWantToBuy := 9
	numberOfPearsWantToBuy := 8
	fruitsCost := float64(numberOfApplesWantToBuy)*priceOfOneApple + float64(numberOfPearsWantToBuy)*priceOfOnePear
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\nВідповідь -", fruitsCost, "грн.")

	// 2. Скільки груш ми можемо купити?
	numberOfPearsCanBuy := int(totalAmount / priceOfOnePear)
	fmt.Println("2. Скільки груш ми можемо купити?\nВідповідь -", numberOfPearsCanBuy, "шт.")

	// 3. Скільки яблук ми можемо купити?
	numberOfApplesCanBuy := int(totalAmount / priceOfOneApple)
	fmt.Println("3. Скільки яблук ми можемо купити?\nВідповідь -", numberOfApplesCanBuy, "шт.")

	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	numberOfPearsWantBuy := 2
	numberOfApplesWantBuy := 2
	totalBudgetOnByuing := float64(numberOfPearsWantBuy)*priceOfOnePear + float64(numberOfApplesWantBuy)*priceOfOneApple
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?\nВідповідь -", totalBudgetOnByuing >= totalAmount)
}
