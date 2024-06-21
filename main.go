package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	skop := scan() //Запись возвращаемого массива в переменную

	//Проверка на размер массива
	if len(skop) > 3 {
		fmt.Print("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(0)
	}
	if len(skop) < 3 {
		fmt.Print("Выдача паники, так как строка не является математической операцией.")
		os.Exit(0)
	}

	//Оператор
	oper := skop[1]

	//Операнды, если перевод в int успешен то записываются, если string то
	first, errFirst := strconv.Atoi(skop[0])
	second, errSecond := strconv.Atoi(skop[2])

	//Вычисляет если это римские цифры
	if errFirst != nil && errSecond != nil {
		first = convertToArab(skop[0])
		second = convertToArab(skop[2])
		if first < second && oper == "-" {
			fmt.Print("Выдача паники, так как в римской системе нет отрицательных чисел.")
			os.Exit(0)
		} else if calc(first, second, oper) == 0 {
			fmt.Print("Выдача паники, так как в римской системе нет обозначения нуля")
			os.Exit(0)
		} else {
			fmt.Print(convertToRim(calc(first, second, oper)))
		}
	} else if first != 0 && second != 0 { //Вычисляет если это арабские цифры
		fmt.Print(calc(first, second, oper))
	} else { //Выводи панику если два операнда из разных систем
		fmt.Print("Выдача паники, так как используются одновременно разные системы счисления.")
		os.Exit(0)
	}
}

// Функция сканера, считывание строки с и возвращение в виде массива
func scan() []string {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	tot := scanner.Text()
	return strings.Split(tot, " ")
}

// Функция перевод в римские цифры
func convertToRim(supply int) string {
	arab := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
		19: "XIX",
		20: "XX",
	}
	return arab[supply]
}

// Функция перевода в арабские цифры
func convertToArab(supply string) int {
	rim := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return rim[supply]
}

// Калькулятор
func calc(first, second int, oper string) int {
	var res int
	switch oper {
	case "/":
		res = first / second
	case "*":
		res = first * second
	case "+":
		res = first + second
	case "-":
		res = first - second
	}
	return res
}
