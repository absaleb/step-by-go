package main

import (
	"errors"
	"fmt"
)

func main()  {
	dec()
	dec2()
}


func dec(){
	var strArr []string
	er := iter(&strArr)
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("nil error")
	}
	fmt.Println(strArr)

	count := 9999
	limit := 100
	pageCount := count/limit

	fmt.Println(pageCount)
}

func dec2(){
	var strArr []string
	er := iter(&strArr)
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("nil error")
	}
	fmt.Println(strArr)

	count := 199
	limit := 100
	pageCount := count/limit

	fmt.Println(pageCount)
}

func iter(str *[]string) error{
	if len(*str) == 20{
		return errors.New("limit exceed")
	}

	*str = append(*str, "q ")

	fmt.Println(*str)

	er := iter(str)

	return er
}