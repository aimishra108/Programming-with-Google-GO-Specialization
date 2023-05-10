package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strings"
)

func main() {
    fmt.Println("input text:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    err := scanner.Err()
    if err != nil {
        log.Fatal(err)
    }
	value_string :=scanner.Text()
	lowerdata1:= strings.ToLower(value_string)
	if ((strings.Contains(lowerdata1,"a")) && (strings.Contains(lowerdata1,"i")) && (strings.Contains(lowerdata1,"n"))){
		fmt.Println("Found the values in the string!!!!")
		os.Exit(0)
		
	}
	fmt.Println("No characters!!!!")

}

{
    package main

    import(
        "bufio"
        "fmt"
        "log"
        "os"
        "strings"
    )

    func main(){
        fmt.Println("Input txt")
        scanner :=bufio.NewScanner(os.Stdin) 
        scanner.Scan()
        err:=scanner.Err()
        if err!=nil{
            log.Fatal(err)
        }
        value_string:=scanner.Text()
        var x [5] int
        x[0] = 2
        var z [5]=[5]{1,2,3,6,7}
        fmt.Printf(x[0])
    }
}