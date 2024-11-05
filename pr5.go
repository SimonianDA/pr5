//Линейное

//No 1

package main

import (
	"fmt"
	"math"
	"strings"
)
func toDecimal(number string, baseFrom int) (int, error) {
	var result int
	number = strings.ToUpper(number)

	for i, digit := range number {
		var value int
		if digit >= '0' && digit <= '9' {
			value = int(digit - '0')
		} else if digit >= 'A' && digit <= 'Z' {
			value = int(digit-'A') + 10
		} else {
			return 0, fmt.Errorf("недопустимый символ %c в числе", digit)
		}

		if value >= baseFrom {
			return 0, fmt.Errorf("цифра %c недопустима для системы счисления с основанием %d", digit, baseFrom)
		}

		result += value * int(math.Pow(float64(baseFrom), float64(len(number)-1-i)))
	}

	return result, nil
}
func fromDecimal(num int, baseTo int) string {
	if num == 0 {
		return "0"
	}

	digits := ""
	for num > 0 {
		remainder := num % baseTo
		if remainder < 10 {
			digits = string('0'+remainder) + digits
		} else {
			digits = string('A'+remainder-10) + digits
		}
		num /= baseTo
	}
	return digits
}

func main() {
	var number string
	var baseFrom, baseTo int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Print("Введите основание системы счисления (от 2 до 36): ")
	fmt.Scan(&baseFrom)
	fmt.Print("Введите основание системы счисления, в которую нужно перевести (от 2 до 36): ")
	fmt.Scan(&baseTo)
	if baseFrom < 2 || baseFrom > 36 || baseTo < 2 || baseTo > 36 {
		fmt.Println("Основание системы счисления должно быть в диапазоне от 2 до 36.")
		return
	}
	decimalValue, err := toDecimal(number, baseFrom)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	result := fromDecimal(decimalValue, baseTo)
	fmt.Printf("Число %s в системе счисления с основанием %d переводится в: %s\n", number, baseTo, result)
}

//No 2

package main

import (
	"fmt"
	"math"
)

var a, b, c float64

func main() {
	fmt.Println("Введите коэффициент a:")
	fmt.Scanln(&a)

	fmt.Println("Введите коэффициент b:")
	fmt.Scanln(&b)

	fmt.Println("Введите коэффициент c:")
	fmt.Scanln(&c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		root1 := ((-1*b) + math.Sqrt(discriminant)) / (2 * a)
		root2 := ((-1*b) - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Два корня: %.2f и %.2f\n", root1, root2)
	} else if discriminant == 0 {
		root := (-1*b) / (2 * a)
		fmt.Printf("Один корень: %.2f\n", root)
	} else {
		realPart := -1*b / (2 * a)
		imaginaryPart := math.Sqrt(-discriminant) / (2 * a)
		fmt.Printf("Комплексные корни: %.2f + %.2fi и %.2f - %.2fi\n", realPart, imaginaryPart, realPart, imaginaryPart)
	}

}

//No 3

package main

import (
	"fmt"
)

var a [6]int

func main() {
	for i := 0; i <= 5; i++ {
		var inp int
		fmt.Scan(&inp)
		a[i] = inp
	}
	for i := 0; i <= 5; i++ {
		for j := i; j < 5; j++ {
			if a[i] > a[j+1] {
				var c = 0
				c = a[j+1]
				a[j+1] = a[i]
				a[i] = c
			}
		}
	}
	fmt.Println(a)
}

//No 4

package main

import (
	"fmt"
)

var a [6]int
var b [6]int
var ab [12]int

func main() {
	for i := 0; i <= 5; i++ {
		var inp int
		fmt.Scan(&inp)
		a[i] = inp
	}
	for i := 0; i <= 5; i++ {
		for j := i; j < 5; j++ {
			if a[i] > a[j+1] {
				var c = 0
				c = a[j+1]
				a[j+1] = a[i]
				a[i] = c
			}
		}
	}

	for i := 0; i <= 5; i++ {
		var inp int
		fmt.Scan(&inp)
		b[i] = inp
	}
	for i := 0; i <= 5; i++ {
		for j := i; j < 5; j++ {
			if b[i] > b[j+1] {
				var c = 0
				c = b[j+1]
				b[j+1] = b[i]
				b[i] = c
			}
		}
	}
	for i := 0; i <= len(a)-1; i++ {
		ab[i] = a[i]
	}
	var j int = len(a)
	for i := 0; i <= len(b)-1; i++ {
		ab[j] = b[i]
		j = j + 1
	}
	for i := 0; i <= 11; i++ {
		for j := i; j < 11; j++ {
			if ab[i] > ab[j+1] {
				var c = 0
				c = ab[j+1]
				ab[j+1] = ab[i]
				ab[i] = c
			}
		}
	}
	fmt.Println(ab)
}

//No 5

package main

import (
	"fmt"
)

var a string
var b string

func main() {
	fmt.Println("Введите основную строку")
	fmt.Scanln(&a)
	fmt.Println("Введите символы для поиска")
	fmt.Scanln(&b)

	var j int = 0
	var ind int = 0
	var i int = 0
	for i = 0; i < len(a)-1; i++ {
		if a[i] == b[j] {
			j = j + 1
			if j == 1 {
				ind = i
			}
			if j == len(b) {
				fmt.Println(ind)
				break
			}
		} else {
			j = 0
		}
	}
	if j < len(b)-1 {
		fmt.Println("-1")
	}
}

//Условия

//No 6

package main

import (
	"fmt"
)

var a int
var b int
var c string
var res int

func main() {

	fmt.Println("Введите значение 1:")
	fmt.Scan(&a)
	fmt.Println("Введите значение 2:")
	fmt.Scan(&b)
	fmt.Println("Введите символ действия:")
	fmt.Scan(&c)

	switch c {
	case "+":
		res = a + b
		fmt.Println(res)
	case "-":
		res = a - b
		fmt.Println(res)
	case "*":
		res = a * b
		fmt.Println(res)
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!!!")
		} else {
			res = a / b
			fmt.Println(res)
		}
	case "%":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!!!")
		} else {
			res = a % b
			fmt.Println(res)
		}
	default:
		fmt.Println("Вы ввели недопустимый символ")
	}
}

//No 7

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(input string) bool {
	
	var cleaned []rune
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned = append(cleaned, unicode.ToLower(char))
		}
	}
	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}

//No 8

package main

import (
	"fmt"
)

var a1 int
var a2 int
var b1 int
var b2 int
var c1 int
var c2 int

func main() {

	fmt.Println("Введите отрезок 1:")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Println("Введите отрезок 2:")
	fmt.Scan(&b1)
	fmt.Scan(&b2)
	fmt.Println("Введите отрезок 3:")
	fmt.Scan(&c1)
	fmt.Scan(&c2)

	if a2 >= b1 && a2 >= c1 && a1 <= b2 && a1 <= c2 {
		fmt.Println(true)
	} else if b2 >= a1 && b2 >= c1 && b1 <= a2 && b1 <= c2 {
		fmt.Println(true)
	} else if c2 >= b1 && c2 >= a1 && c1 <= b2 && c1 <= a2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

//No 9

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func findLongestWord(sentence string) string {
	re := regexp.MustCompile(`[^\w\s]+`)
	cleaned := re.ReplaceAllString(sentence, "")
	words := strings.Fields(cleaned)
	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	return longestWord
}

func main() {
	var input string
	fmt.Print("Введите предложение: ")
	fmt.Scanln(&input)
	longestWord := findLongestWord(input)
	if longestWord != "" {
		fmt.Println("Самое длинное слово:", longestWord)
	} else {
		fmt.Println("Не удалось найти слова в предложении.")
	}
}

//No 10

ackage main

import (
	"fmt"
)

var year int
var a, b, c int

func main() {

	fmt.Println("Введите год:")
	fmt.Scan(&year)

	a = year % 4
	b = year % 100
	c = year % 400
	if a == 0 && b != 0 {
		fmt.Print(true)
	} else if c == 0 {
		fmt.Print(true)
	} else {
		fmt.Print(false)
	}

}

//Циклы

//No 11

package main

import (
	"fmt"
)

var limit int

func main() {

	fmt.Print("Введите верхнюю границу для чисел Фибоначчи: ")
	fmt.Scan(&limit)

	a, b := 0, 1

	for a <= limit {
		fmt.Println(a)
		a, b = b, a+b
	}
}


//No 12

package main

import (
	"fmt"
)

var a, b int

func main() {

	fmt.Println("Введите числа:")
	fmt.Scan(&a, &b)

	for a = a + 1; a <= b-1; a++ {
		fmt.Println(a)
	}
}

//No 13

package main

import (
	"fmt"
	"math"
)
func isArmstrong(number int) bool {
	digits := make([]int, 0)
	n := number
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	numDigits := len(digits)
	sum := 0
	for _, digit := range digits {
		sum += int(math.Pow(float64(digit), float64(numDigits)))
	}
	return sum == number
}

func main() {
	var lower, upper int
	fmt.Print("Введите нижнюю границу диапазона: ")
	fmt.Scan(&lower)
	fmt.Print("Введите верхнюю границу диапазона: ")
	fmt.Scan(&upper)
	fmt.Printf("Числа Армстронга в диапазоне от %d до %d:\n", lower, upper)
	for i := lower; i <= upper; i++ {
		if isArmstrong(i) {
			fmt.Println(i)
		}
	}
}

//No 14

package main

import (
	"fmt"
)

func reverseString(s string) string {

	reversed := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		reversed[len(s)-1-i] = s[i]
	}

	return string(reversed)
}

func main() {
	var input string

	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	reversedString := reverseString(input)

	fmt.Printf("Обратная строка: %s\n", reversedString)
}

//No 15

package main

import (
	"fmt"
)
func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func main() {
	var num1, num2 int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)
	result := gcd(num1, num2)
	fmt.Printf("Наибольший общий делитель чисел %d и %d равен %d\n", num1, num2, result)
}