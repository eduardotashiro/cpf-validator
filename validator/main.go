package main

import (
	"fmt"
	"regexp"
)

func validateFirstDigit(cpf string) {
	cleanCpf := regexp.MustCompile(`[^\d]+`).ReplaceAllString(cpf, "")
	fmt.Println(cleanCpf)
	sum := 0
	for i := 0; i < len(cleanCpf); i++ {
		fmt.Printf("%c \n", cleanCpf[i])
	}

	println(sum)
}

func main() {
	validateFirstDigit("529.982.247-25")
}

// \d
// \w
// \s
// []
// [^]
// +

// func validatorFistDigit(cpf string) string {
// 	// cpf = "529.982.247-25"
// 	// fmt.Println(cpf)
// 	// 	var sum int
// 	// 	println(sum)

// 	//	for i := 0; i < 9; i++ {
// 	//		// result :=
// 	//	}
// 	//
// 	// //    sum += result * (10 - i)
// 	//
// 	//	// (sum * 10) % 11
// 	fmt.Println(cpf)
// 	return cpf

// }
