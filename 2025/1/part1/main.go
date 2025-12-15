package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	count := 0

	position:=50

	file, err := os.Open("input.txt")
	if err!=nil {
		fmt.Println("Could not read file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line:= scanner.Text()

		direction := line[0]
		value, err:= strconv.Atoi(line[1:])

		if err!=nil {
			fmt.Println("Failed to get value")
			return
		}

		if direction == 'R'{
			position = (position + value)%100
		} else{
			position = (position - value + 100)%100
		}

		if position == 0 {count++;}

	     	
	}

	fmt.Printf("Password is : %d", count)

}