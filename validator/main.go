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
// 	if repCPF[0] == repCPF[1] && repCPF[1] == repCPF[2] && repCPF[2] == repCPF[3] && repCPF[3] == repCPF[4] && repCPF[4] == repCPF[5] && repCPF[5] == repCPF[6] && repCPF[6] == repCPF[7] && repCPF[7] == repCPF[8] && repCPF[8] == repCPF[9] && repCPF[9] == repCPF[10] && repCPF[10] == repCPF[0] {
// 		fmt.Printf("cpf : %s , não pode repetir os números !!", repCPF)
// 	}
// 	return repCPF
// }

func validateCPF(c string) bool {
	cpf := removeCharacters(c)
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

	rest1 := (sum * 10) % 11
	if rest1 < 10 {
		return rest1 == int(cpf[9])
	}
	return int(cpf[9] == 0)

}

func main() {
	//  529 982 247 25
	cpf := "529.982.247-25"
	validateCPF(cpf)
}
