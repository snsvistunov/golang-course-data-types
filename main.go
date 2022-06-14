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

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	const priceOfOneApple = 5.99
	const priceOfOnePear = 7.00
	const currency = "грн."
	const totalAmount = 23.00
	const firstQuestion = "1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?"
	const secondQuestion = "2. Скільки груш ми можемо купити?"
	const thirdQuestion = "3. Скільки яблук ми можемо купити?"
	const fourthQuestion = "4. Чи ми можемо купити 2 груші та 2 яблука?"
	const firstAnswer = "Вам потрібно"
	const youCanBuyAnswerPart = "Ви можете придбати "
	const secondAnswerVariantOne = " грушу"
	const secondAnswerVariantTwo = " груші"
	const secondAnswerVariantThree = " груш"
	const degreeOfRounding = 100 //ступінь округлення до двох знаків
	const youDontHaveEnoughMoney = "У вас недостатьно грошей!"

	var numberOfPearsCanBuy float64
	// var numberOfApplesCanBuy byte
	var numberOfPearsWantBuy float64
	var numberOfApplesWantBuy float64
	var budgetToBuyApllesAndPears float64
	var secondAnswer string

	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?

	numberOfApplesWantBuy = 9
	numberOfPearsWantBuy = 8
	fmt.Println(firstQuestion)
	budgetToBuyApllesAndPears = numberOfApplesWantBuy*priceOfOneApple + numberOfPearsWantBuy*priceOfOnePear //до округлення
	budgetToBuyApllesAndPears = math.Round(budgetToBuyApllesAndPears*degreeOfRounding) / degreeOfRounding   //округлення до найближчого
	fmt.Println(firstAnswer, budgetToBuyApllesAndPears, currency)

	// 2. Скільки груш ми можемо купити?
	fmt.Println(secondQuestion)
	numberOfPearsCanBuy = math.Trunc(totalAmount / priceOfOnePear)
	if numberOfPearsCanBuy < 1 {
		secondAnswer = youDontHaveEnoughMoney
	} else if numberOfPearsCanBuy < 5 {
		secondAnswer = youCanBuyAnswerPart + strconv.FormatFloat(numberOfPearsCanBuy, 'f', 0, 64) + secondAnswerVariantTwo
	} else {
		secondAnswer = youCanBuyAnswerPart + strconv.FormatFloat(numberOfPearsCanBuy, 'f', 0, 64) + secondAnswerVariantThree
	}

	fmt.Println(secondAnswer)
}
