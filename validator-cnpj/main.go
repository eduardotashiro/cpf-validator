package main

import (
	"fmt"
	"regexp"
	// "strconv"
)

var (
	cnpjd1 string
	cnpjd2 string
	cnpjP1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjP2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

func removeCharacters(cleanCNPJ string) string {
	re := regexp.MustCompile(`[^\d]+`)
	cnpj := re.ReplaceAllString(cleanCNPJ, "")
	return cnpj
}

func validateCnpj(c string) bool {
	cnpj := removeCharacters(c)

	if len(cnpj) != 14 {
		return false
	}

	counter := 0
	for k := 0; k < len(cnpj); k++ {
		if cnpj[0] == cnpj[k] {
			counter++
		}
		if counter == len(cnpj) {
			return false
		}
	}

	//digito verificador
	p1 := cnpj[12:]
	p2 := cnpj[:12]
	fmt.Println(p1, p2)

	return true
}

func main() {
	var CNPJ string = "59.541.264/0001-03"
	var CNPJ2 string = "11.111.111/1111-11"
	test := validateCnpj(CNPJ)
	test2 := validateCnpj(CNPJ2)
	fmt.Println(test, test2)
}
