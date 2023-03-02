package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var romanNumbers = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
var arabicNumbers = map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}

func checkIsNumber(in string) bool {
	r := []rune(in)
	for _, e := range r {
		if !unicode.IsNumber(e) {
			return false
		}
	}
	n, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	if n < 1 && n > 10 {
		return false
	} else {
		return true
	}
}
func checkIsRoman(in string) bool {
	for _, e := range arabicNumbers {
		in = strings.ReplaceAll(in, e, "")
	}
	if in == "" {
		return true
	} else {
		return false
	}
}

func romanToArabic(in string) (int, error) {
	r := []rune(in)
	c := 0
	ex, prev := "", ""
	count := 0
	var err error
	for i := 0; i < len(r); i++ {
		if ex == string(r[i]) {
			err = fmt.Errorf("Ошибка в арабском числе.")
			return c, err
		}
		if i+1 < len(r) && romanNumbers[string(r[i])] < romanNumbers[string(r[i+1])] {
			c += romanNumbers[string(r[i+1])] - romanNumbers[string(r[i])]
			ex = string(r[i+1])
			i++
		} else {
			if ex != "" && prev == string(r[i]) {
				err = fmt.Errorf("Ошибка в арабском числе.")
				return c, err
			} else {
				ex = ""
			}
			if prev != string(r[i]) {
				prev = string(r[i])
				count = 1
			} else {
				count++
				if count > 3 {
					err = fmt.Errorf("Ошибка в арабском числе.")
					return c, err
				}
			}
			c += romanNumbers[string(r[i])]
		}
	}
	return c, nil
}

func arabicToRoman(in int) string {
	r := []rune(strconv.Itoa(in))
	out := ""
	for i, e := range r {
		dis := len(r) - i - 1
		rome := int(math.Pow10(dis))
		sw, _ := strconv.Atoi(string(e))
		switch {
		case sw < 4:
			for j := 1; j <= sw; j++ {
				out += arabicNumbers[rome]
			}
		case sw == 4:
			out += arabicNumbers[rome] + arabicNumbers[rome*5]
		case sw < 9:
			out += arabicNumbers[rome*5]
			for j := 6; j <= sw; j++ {
				out += arabicNumbers[rome]
			}
		case sw == 9:
			out += arabicNumbers[rome] + arabicNumbers[rome*10]
		}
	}
	return out
}

func calculate(a int, b int, o string) int {
	switch o {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	default:
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var a, b int
	var err error
	separator := []string{"+", "-", "/", "*"}
	fmt.Fprintf(writer, "%s ", "Добро пожаловать в калькулятор!\n Калькулятор умеет складывать, вычитать, умножать и делить два целых числа от 1 до 10. \n Числа могут быть как в арабском формате так и в римском. \n Введите арифметическое действие. \n (для выхода введите exit): \n")
	writer.Flush()
	for {
		line, _ := reader.ReadString('\n')
		line = strings.ReplaceAll(strings.ReplaceAll(line, " ", ""), "\r\n", "")
		line = strings.ToUpper(line)
		if line == "EXIT" {
			break
		}
		isArabic := true
		calculated := false
		for _, e := range separator {
			splitted := strings.Split(line, e)
			if len(splitted) == 2 {
				calculated = true
				if checkIsRoman(splitted[0]) && checkIsRoman(splitted[1]) {
					isArabic = false
					a, err = romanToArabic(splitted[0])
					if err != nil {
						fmt.Fprintf(writer, "%s \n", err)
						break
					}
					b, err = romanToArabic(splitted[1])
					if err != nil {
						fmt.Fprintf(writer, "%s \n", err)
						break
					}
				} else if checkIsNumber(splitted[0]) && checkIsNumber(splitted[1]) {
					isArabic = true
					a, err = strconv.Atoi(splitted[0])
					if err != nil {
						fmt.Fprintf(writer, "%s \n", err)
						break
					}
					b, err = strconv.Atoi(splitted[1])
					if err != nil {
						fmt.Fprintf(writer, "%s \n", err)
						break
					}
				} else {
					fmt.Fprintf(writer, "%s \n", "Не верный формат вводных данных.")
					break
				}
				if a < 1 || a > 10 || b < 1 || b > 10 {
					fmt.Fprintf(writer, "%s \n", "Значения выходят за предел 1 - 10.")
					break
				}
				res := calculate(a, b, e)
				if !isArabic {
					if res < 0 {
						fmt.Fprintf(writer, "%s \n", "В римской системе исчисления нет отрицательных значений.")
						break
					}
					fmt.Fprintf(writer, "%s \n", arabicToRoman(res))
				} else {
					fmt.Fprintf(writer, "%d \n", res)
				}
				break
			}
		}
		if !calculated {
			fmt.Fprintf(writer, "%s \n", "Не верный формат вводных данных.")
		}
		writer.Flush()
	}
}
