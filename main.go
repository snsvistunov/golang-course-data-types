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

	const currency = "грн."

	const firstQuestion = "1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?"
	const secondQuestion = "2. Скільки груш ми можемо купити?"
	const thirdQuestion = "3. Скільки яблук ми можемо купити?"
	const fourthQuestion = "4. Чи ми можемо купити 2 груші та 2 яблука?"
	const firstAnswer = "Вам потрібно"
	const youCanBuyAnswerPart = "Ви можете придбати "
	const secondAnswerVariantOne = " грушу"
	const secondAnswerVariantTwo = " груші"
	const secondAnswerVariantThree = " груш"
	const thirdAnswerVariantOne = " яблуко"
	const thirdAnswerVariantTwo = " яблука"
	const thirdAnswerVariantThree = " яблук"
	const degreeOfRounding = 100 //ступінь округлення до двох знаків
	const youDontHaveEnoughMoney = "Ні. У вас недостатьно грошей!"
	const youCantBuyFruits = "У вас недостатьно грошей!"
	const youHaveEnoughMoney = "Так. У вас достатьно грошей!"

	var priceOfOneApple = 5.99
	var priceOfOnePear = 7.00
	var totalAmount = 23.00

	var numberOfPearsCanBuy int64
	var numberOfApplesCanBuy int64
	var numberOfPearsWantBuy float64
	var numberOfApplesWantBuy float64
	var budgetToBuyApllesAndPears float64
	var secondAnswer string
	var thirdAnswer string
	var totalBudget float64

	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?

	numberOfApplesWantBuy = 9
	numberOfPearsWantBuy = 8
	fmt.Println(firstQuestion)
	budgetToBuyApllesAndPears = numberOfApplesWantBuy*priceOfOneApple + numberOfPearsWantBuy*priceOfOnePear //до округлення
	budgetToBuyApllesAndPears = math.Round(budgetToBuyApllesAndPears*degreeOfRounding) / degreeOfRounding   //округлення до найближчого
	fmt.Println(firstAnswer, budgetToBuyApllesAndPears, currency)

	// 2. Скільки груш ми можемо купити?
	fmt.Println(secondQuestion)
	numberOfPearsCanBuy = int64(math.Trunc(totalAmount / priceOfOnePear))
	if numberOfPearsCanBuy%10 < 1 {
		secondAnswer = youCantBuyFruits
	} else if numberOfPearsCanBuy%10 == 1 {
		secondAnswer = youCanBuyAnswerPart + strconv.FormatInt(numberOfPearsCanBuy, 10) + secondAnswerVariantOne
	} else if numberOfPearsCanBuy%10 < 5 {
		secondAnswer = youCanBuyAnswerPart + strconv.FormatInt(numberOfPearsCanBuy, 10) + secondAnswerVariantTwo
	} else {
		secondAnswer = youCanBuyAnswerPart + strconv.FormatInt(numberOfPearsCanBuy, 10) + secondAnswerVariantThree
	}
	fmt.Println(secondAnswer)

	// 3. Скільки яблук ми можемо купити?
	fmt.Println(thirdQuestion)
	numberOfApplesCanBuy = int64(math.Trunc(totalAmount / priceOfOneApple))
	if numberOfApplesCanBuy%10 < 1 {
		secondAnswer = youCantBuyFruits
	} else if numberOfApplesCanBuy%10 == 1 {
		secondAnswer = youCanBuyAnswerPart + strconv.FormatInt(numberOfPearsCanBuy, 10) + thirdAnswerVariantOne
	} else if numberOfApplesCanBuy%10 < 5 {
		thirdAnswer = youCanBuyAnswerPart + strconv.FormatInt(numberOfApplesCanBuy, 10) + thirdAnswerVariantTwo
	} else {
		thirdAnswer = youCanBuyAnswerPart + strconv.FormatInt(numberOfApplesCanBuy, 10) + thirdAnswerVariantThree
	}

	fmt.Println(thirdAnswer)

	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	fmt.Println(fourthQuestion)
	numberOfPearsWantBuy = 2
	numberOfApplesWantBuy = 2
	totalBudget = numberOfPearsWantBuy*priceOfOnePear + numberOfApplesWantBuy*priceOfOneApple
	if totalBudget > totalAmount {
		fmt.Println(youDontHaveEnoughMoney)
	} else {
		fmt.Println(youHaveEnoughMoney)
	}
}
