package main

import (
	"fmt"
)

var (
	max int8 = 14 // digitos mais os caracteres
	div int8 = 11 // digitos
)

func calculaDigitos(cpf []int, a int) {
	var (
		b      int = a - 1
		i      int
		dig    int 
		soma   int 
		result int 
		rest   int 
	)

	for i = 0; i < b; i++ {
		
	}
}

func validador(a int, b int, c[]int){


}



func main() {

}

// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// )

// func validaCpf(cpf string) bool {
// 	re := regexp.MustCompile(`[^\d]+`)
// 	cpf = re.ReplaceAllString(cpf, "")
// 	if len(cpf) != 11 {
// 		return false
// 	}

// 	sum := 0
// 	for i := range 9 {
// 		digit, err := strconv.Atoi(string(cpf[i]))
// 		if err != nil {
// 			return false
// 		}
// 		sum += digit * (10 - i)
// 	}
// 	r := (sum * 10) % 11
// 	string(r)
// 	if r == cpf[9]{

// 	}

// 	// resto := (sum * 10) % 11
// 	fmt.Println(sum) //295 :)
// 	// }
// 	// resto := (sum * 10) % 11
// 	// if resto < 10 {
// 	// 	return int(cleanCpf[9]) == resto
// 	// }
// 	return true
// 	//int(cleanCpf[9]) == 0
// }

// func main() {
// 	//  529 982 247 25
// 	cpf := "529.982.247-25"
// 	validaCpf(cpf)
// }
