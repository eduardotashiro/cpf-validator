package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func removeCharacters(cleanCPF string) string {
	re := regexp.MustCompile(`[^\d]+`)
	cpf := re.ReplaceAllString(cleanCPF, "")
	return cpf
}

// func isRepetitive(repCPF string) string {
// 	if len(repCPF) == 0 {
// 		fmt.Printf("cpf : %s , vazio !!", repCPF)
// 	}

// 	return repCPF
// }

func validateCPF(c string) bool {
	//remove caracteres como .(ponto) e -(traço)
	cpf := removeCharacters(c)

	//verifica se o CPF digitado tem 11 digitos, pois ja removi . -
	if len(cpf) != 11 {
		return false
	}

	//verifica se todos os dígitos são iguais
	if cpf[0] == cpf[1] && cpf[1] == cpf[2] && cpf[2] == cpf[3] && cpf[3] == cpf[4] && cpf[4] == cpf[5] && cpf[5] == cpf[6] && cpf[6] == cpf[7] && cpf[7] == cpf[8] && cpf[8] == cpf[9] && cpf[9] == cpf[10] && cpf[10] == cpf[0] {
		return false
	}

	//calculando primeiro dígito verificador
	sum := 0
	for i := range 9 {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += digit * (10 - i)
	}

	rest := (sum * 10) % 11
	if rest < 10 {
		return rest == int(cpf[9])
	}
	return int(cpf[9] == 0)

}

func main() {
	//  529 982 247 25
	cpf := "529.982.247-25"
	validateCPF(cpf)
}
