package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func validaCpf(cpf string) bool {
	re := regexp.MustCompile(`[^\d]+`)
	cpf = re.ReplaceAllString(cpf, "")
	if len(cpf) != 11 {
		return false
	}
	sum := 0
	for i := range 9 {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += digit * (10 - i)
	}

	// resto := (sum * 10) % 11
	fmt.Println(sum)
	// }
	// resto := (sum * 10) % 11
	// if resto < 10 {
	// 	return int(cleanCpf[9]) == resto
	// }
	return true
	//int(cleanCpf[9]) == 0
}

func main() {
	//  529 982 247 25
	cpf := "529.982.247-25"
	validaCpf(cpf)
}
