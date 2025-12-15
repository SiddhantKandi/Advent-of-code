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

		mod_value := value%100;
		count+=value/100
		if direction == 'R' {
			position+=mod_value
			if position > 99{
				count++
				position = position - 100
			}
		} else{
			old_position := position
			position-=mod_value
			if position < 0 {
				if(old_position !=0) {count++;}
				position +=100
			}

			if position == 0 {count++;}

		}
	}

	fmt.Printf("Password is : %d",count)

}