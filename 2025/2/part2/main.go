package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Could not read the input file", err)
		return
	}

	defer file.Close()


	scanner := bufio.NewScanner(file)

	count := int64(0)

	for scanner.Scan(){
		line := scanner.Text()

		parts := strings.Split(line, "-")

		if len(parts) != 2 {
			fmt.Println("Error splitting the string")
			return
		}

		start, err := strconv.ParseInt(parts[0], 10, 64)

		if err!=nil{
			fmt.Println("Invalid number")
			return
		}

		end, err := strconv.ParseInt(parts[1], 10, 64)

		if err!=nil{
			fmt.Println("Invalid number")
			return
		}


		for start<=end {
		  startStr := strconv.FormatInt(start, 10)

		  flag := false

		  n:= len(startStr)

		  for k:=1;k<=n/2;k++{

			if n%k != 0 {
				continue;
			}

			pattern:= startStr[:k]

			valid := true

			for i:=k;i<n;i+=k{
				if startStr[i:i+k]!=pattern{
					valid = false
					break;
				}
			}

			if valid{
				flag = true
				break
			}

		  }

		  if flag {
			fmt.Println("Number is : ",start)
			count+=start
		  }

		  start++
		}
		
	}
	fmt.Println("Count is : ",count)
}