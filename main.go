package main

import "fmt"

func romanToInt(s string) {
	var result int
	integer := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i, v := range s {
		if i < len(s)-1 {
			if v == 'I' && (s[i+1] == 'V' || s[i+1] == 'X') {
				result -= 1
				continue
			}
			if v == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
				result -= 10
				continue
			}
			if v == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
				result -= 100
				continue
			}
		}
		result += integer[uint8(v)]
	}
	fmt.Println(result)
}

func main() {
	var s string
	for {
		fmt.Println("Введите римское число заглавными латинскими буквами или exit для завершения программы")
		fmt.Scanln(&s)
		if s == "exit" {
			break
		}
		romanToInt(s)
	}
}
