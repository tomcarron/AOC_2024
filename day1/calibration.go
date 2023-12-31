package main

import ( 
	"fmt"
	"unicode/utf8"
	"strconv"
	"log"
	"os"
	"bufio"
	"strings"
	"regexp"
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

func part1() {
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
	fmt.Println("part 1",total)
}


func calc_val_2(s string) (value string) {	
	words :=[]string{"zero","one","two","three","four","five","six","seven","eight","nine"}
	nums :=[]string{"0","1","2","3","4","5","6","7","8","9"}
	// Create a regex pattern 
	pattern := strings.Join(append(nums, words...), "|") //appending nums with words (... unpacks words so theyre appended one by one) Join concatenates the elements of its first argument to create a single string. (with "|" as seperator)
	regex := regexp.MustCompile(pattern)		//makes a regex pattern to search with

	// Find matches in s
	matches := regex.FindAllString(s, -1) 

	// Extract the answers
	firstValue := matches[0]
	lastValue := matches[len(matches)-1]

	// Map words to nums
	wordToNum := make(map[string]string)
	for i, word := range words {
		wordToNum[word] = nums[i]
	}

	// Replace words with nums only if not already in num form
	if num, exists := wordToNum[firstValue]; exists {
		firstValue = num
	}
	if num, exists := wordToNum[lastValue]; exists {
		lastValue = num
	}


	var result string = firstValue+lastValue
	return result
}



func part2() {
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
		num, err := strconv.Atoi(calc_val_2(scanner.Text()))
		fmt.Println(scanner.Text(),num)

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
	fmt.Println("part 2",total)
}

func main() {
	//part1()

	part2()

 	//const test string = "xftwo3af4nine0"
	//fmt.Println(test)
	//fmt.Println(first_val(test))
	//fmt.Println(reverse_str(test))
	//fmt.Println(second_val(test))
	//fmt.Println(calc_value(test))
	//fmt.Println(calc_val_2(test)) 
	
}