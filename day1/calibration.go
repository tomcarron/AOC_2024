package main

import ( 
	"fmt"
	"unicode/utf8"
	"strconv"
	"log"
	"os"
	"bufio"
)
// get first rune, check if int if no go to second
func calc_value(s string) string {
	a:=strconv.Itoa(first_val(s))
	b:=strconv.Itoa(second_val(s))
	var result string = a+b
	return result
}

func first_val(s string) int {	
	var firstval int 
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		//fmt.Println(runeValue)
		if runeValue == '0' || runeValue == '1' || runeValue ==  '2' || runeValue ==  '3' || runeValue ==  '4' || runeValue ==  '5' || runeValue ==  '6' || runeValue ==  '7' || runeValue ==  '8' || runeValue ==  '9' {
			firstval = int(runeValue - '0')
			break
		}
		w = width
	}	
	return firstval
}

func reverse_str(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func second_val(s string) int {
	//reverse s
	s2:=reverse_str(s)
	return first_val(s2)
}
func main() {
	// open the file using Open() function from os library
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	var total int = 0
	for scanner.Scan() {
		//Print the line
		//fmt.Println(scanner.Text())
		num, err := strconv.Atoi(calc_value(scanner.Text()))

		if err != nil {
			// Handle the error if the conversion fails
			fmt.Println("Error:", err)
			return
		}
	
		total = total + num
	}
    // check for the error that occurred during the scanning
    
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    
	// Close the file at the end of the program
	file.Close()
	fmt.Println(total)

/* 	const test string = "a1th6k3l"
	fmt.Println(test)
	fmt.Println(first_val(test))
	fmt.Println(reverse_str(test))
	fmt.Println(second_val(test))
	fmt.Println(calc_value(test)) */
	
}