package main

import(
	"os"
	"fmt"
	"strings"
	"bufio"
	"strconv"
)

func ReadConsole(message string)(string,error){
	fmt.Println(message)
	scanner:=bufio.NewReader(os.Stdin)
	data,_,err:=scanner.ReadLine()
	// ns:=strings.Split(string(data)," ")
	// for _,s:=range ns{
	// 	n,_:=strconv.Atoi(s)
	// 	values=append(values,n)
	// }
	return string(data),err

}

func RemoveSpecialReturnSymbols(s string) string {
	var trailing = "\n"

	str := strings.ReplaceAll(s, trailing, "") // remove new line

	return str
}

func insertionSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		tmp := numbers[i] 
		j := i - 1        
		for (j >= 0) && (numbers[j] >= tmp) {
			numbers[j+1] = numbers[j] 
			j = j - 1                 
		}
		numbers[j+1] = tmp
	}
}

func main() {
	slice := make([]int, 0)
	var input string
	var counter int

	for (input != "X" || input != "x"){
		if input, err := ReadConsole("Insert an integer or X to exit: "); err != nil {
			fmt.Println("Input error:", err, ".Please try again.")
			continue
		} else {
			fmt.Println("Input string data:", input)
			input = RemoveSpecialReturnSymbols(input)
			if (input == "X" || input == "x") {
				fmt.Println("X or x entered, exit")
				break
			}

			if number, err := strconv.Atoi(input); err != nil {
				fmt.Println("Input string is not an integer number...")
				continue
			} else {
				fmt.Println("Input string is an integer number: ", number)

				// if counter < 3 {
				// 	slice[counter] = number
				// 	fmt.Println("Slice not sorted:", slice)
				// 	counter++ 
				// 	continue
				// }
				slice = append(slice, number)
				insertionSort(slice)
				fmt.Println("Slice in the (increasing) order:", slice)

				counter++ 
			}
		}
	}
}
