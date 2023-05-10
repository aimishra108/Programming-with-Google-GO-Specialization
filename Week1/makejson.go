package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person map[string] string

func RemoveSuffixes(s string)string {
	var newline="\n"
	var cabaret="\r"
	if strings.HasSuffix(s,newline){
		s=strings.TrimSuffix(s, newline)
	}
	if strings.HasSuffix(s,cabaret){
		s=strings.TrimSuffix(s, cabaret)
	}
	return s
}

func GetPersonData() Person{ 
	person:=make(Person)
	stdinReader:=bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Enter the name: ")
		if name_data,err:=stdinReader.ReadString('\n'); err!=nil {
			fmt.Println("\nError",err)
		} else {
			person["name"]=RemoveSuffixes(name_data)
			break
			
		}
	}

	for{
		fmt.Println("Enter the address: ")
		if add_data,err:=stdinReader.ReadString('\n');err!=nil{
			fmt.Println("\nError",err)
		} else {
			person["address"]=RemoveSuffixes(add_data)
			break
		}	
	}
	return person
}

func PersonJSON(person Person){
	if json_data,err:=json.Marshal(person); err!= nil {
		fmt.Println("\nError parsing in JSON")

	} else {
		fmt.Println("\nJSON data",string(json_data))

	}
	

} 

func main(){
	person:=GetPersonData()
	PersonJSON(person)

}