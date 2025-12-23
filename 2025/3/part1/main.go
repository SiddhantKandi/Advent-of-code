package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Could not read the input file", err)
		return
	}

	defer file.Close()


	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan(){
		line := scanner.Text()	
		
		maxi:= 0
		maxi2 := 0

		index:=0

		for i:=0;i<len(line)-1;i++{
			if int(line[i]-'0') > maxi {
				index = i
				maxi = int(line[i]-'0') 
			}
		}

		for i:=index+1;i<len(line);i++{
			if int(line[i]-'0') > maxi2 {
				maxi2 = int(line[i]-'0') 
			}
		}

		count += maxi*10 + maxi2
		
	}
	fmt.Println("Count is : ",count)
}