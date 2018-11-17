package main

import (
	"errors"
	"fmt"
	"gitlab.okta-solutions.com/mashroom/backend/verification"
	"gitlab.okta-solutions.com/mashroom/backend/verification/impl"
)

func main() {
	dec()
}

func dec() {
	var field verification.VerifyPostcodeQueryRequest
	field.QueryLine = "Christchurch Fairmile Road"

	resp, err := impl.VerifyPostcodeQueryImpl(&field)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(resp.City)
		fmt.Println(len(resp.AddressLine))
		fmt.Println(resp.AddressLine)
	}
}

func iter(str *[]string) error {
	if len(*str) == 20 {
		return errors.New("limit exceed")
	}

	*str = append(*str, "q ")

	fmt.Println(*str)

	er := iter(str)

	return er
}
