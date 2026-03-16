package main

import (
	
	"regexp"
)

func validaCpf(cpf string) bool {
	cleanCpf := regexp.MustCompile(`[^\d]+`).ReplaceAllString(cpf, "")

	if len(cleanCpf) != 11 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cleanCpf[i]) * (10 - i)
		// resto := (sum * 10) % 11
		// fmt.Println(sum)
	}

	resto := (sum * 10) % 11

	if resto < 10 {
		return int(cleanCpf[9]) == resto
	}

	return int(cleanCpf[9]) == 0
}

func main() {
	//  529 982 247 25
	cpf := "529.982.247-25"
	validaCpf(cpf)
}
