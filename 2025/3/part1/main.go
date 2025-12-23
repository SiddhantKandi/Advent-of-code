package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMaxi(s string, num *int64,index *int){

	maxi := int64(0)

	for i:=0;i<len(s);i++ {
		if maxi < int64(s[i]-'0') {
			maxi = int64(s[i]-'0')
			*index = i
		}
	}

	*num = (*num)*10 + maxi
}

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
		num := int64(0)

		index := 0

		start :=0

		n:= len(line)

		for i:=range 12{
			findMaxi(line[start:n-11+i],&num,&index)
			start += index+1
		}

		count += num
		
	}
	fmt.Println("Count is : ",count)
}