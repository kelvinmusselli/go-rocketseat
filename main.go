package main

import "fmt"

func main(){
	// fmt.Println("Hello World")
	// fmt.Println(pacote.Foo)
	// pacote.PrintMinha()

	// fmt.Println(somar(1,2))
	// a, b := swap(10,20)
	// fmt.Println(a,b)

	// fmt.Println(dividir(10,2))
	// x := somarHighOrder(2)(3)
	// f:= somarHighOrder(1)
	// x:= f(3)
	// fmt.Println(x)

	fmt.Println(somarvariadas(1,2,3))
}

//se a primeira letra da funcao ou var for
//minuscula é privada se for mais maiuscula
// é publica

// func somar(a int, b int) int {
// 	return a+b
// }

// func swap (a,b int) (int, int) {
// 	return b,a
// }

// func dividir(a, b int) (res int, rem int) {
// 	res = a/b
// 	rem = a%b
// 	return res,rem
// }

// func somarHighOrder(a int) func(int) int {
// 	return func(b int) int {
// 		return a+b
// 	}
// }

//funcao variatica
func somarvariadas (nums ...int) int {
	var out int
	for _, n:= range nums {
		out += n
	}
	return out
}