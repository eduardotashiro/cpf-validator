package main

import (
	"fmt"
	"regexp"
	// "strconv"
)

func removeCharacters(cleanCNPJ string) string {
	re := regexp.MustCompile(`[^\d]+`)
	cnpj := re.ReplaceAllString(cleanCNPJ, "")
	return cnpj
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func validateCnpj(c string) bool {
	
	return true
}

func main() {

	var CNPJ string = "59.541.264/0001-03"
	cnpjInv := reverse(CNPJ)
	fmt.Println(cnpjInv)

}
