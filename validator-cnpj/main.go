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
	cnpj := removeCharacters(c)

	if len(cnpj) != 14 {
		return false
	}

	if  cnpj[0]  == cnpj[1] &&
		cnpj[1]  == cnpj[2] &&
		cnpj[2]  == cnpj[3] &&
		cnpj[3]  == cnpj[4] &&
		cnpj[4]  == cnpj[5] &&
		cnpj[5]  == cnpj[6] &&
		cnpj[6]  == cnpj[7] &&
		cnpj[7]  == cnpj[8] &&
		cnpj[8]  == cnpj[9] &&
		cnpj[9]  == cnpj[10] &&
		cnpj[10] == cnpj[11] &&
		cnpj[11] == cnpj[12] &&
		cnpj[12] == cnpj[13] &&
		cnpj[13] == cnpj[0] {
		return false
	}

	// calcular aqui 

	return true
}

func main() {

	var CNPJ string = "59.541.264/0001-03"
	fmt.Println(CNPJ)
	cnpjInv := reverse(CNPJ)
	fmt.Println(cnpjInv)

}
