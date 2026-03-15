package main

import (
	"fmt"
	"regexp"
)

func validaCpf(cpf string) bool {
	re := regexp.MustCompile(`[^\d]+`)
	cpf = re.ReplaceAllString(cpf, "")
	fmt.Println(cpf)
	return false
}

func main() {
	    //  529 982 247 25
	cpf := "529.982.247-25" 
	validaCpf(cpf)
}
