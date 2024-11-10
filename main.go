package main

import "fmt"

// func main(){
// 	// fmt.Println("Hello World")
// 	// fmt.Println(pacote.Foo)
// 	// pacote.PrintMinha()

// 	// fmt.Println(somar(1,2))
// 	// a, b := swap(10,20)
// 	// fmt.Println(a,b)

// 	// fmt.Println(dividir(10,2))
// 	// x := somarHighOrder(2)(3)
// 	// f:= somarHighOrder(1)
// 	// x:= f(3)
// 	// fmt.Println(x)

// 	// fmt.Println(somarvariadas(1,2,3))

// 	// var nome, sobrenome string = "Kelvin", "Musselli"

// 	// fmt.Println(nome, sobrenome)
// }

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
// func somarvariadas (nums ...int) int {
// 	var out int
// 	for _, n:= range nums {
// 		out += n
// 	}
// 	return out
// }

// converter um valor para outro
// func main() {
// 	var x int = 65
// 	f:= float64(x)
// }

//arrays

// func main() {
// 	// como especificar um valor para uma posicao do array
// 	arr := [10]int{5:400, 0:1}
// 	fmt.Println(arr)
// }

func main() {
	var result int

	for i := 0; i < 10; i++ {
		result++
	}

	// var i int
	// for ; i < 10;  {
	// 	result++
	// }


	arr :=[10]int{1,2,3,4,5,6,7,8,9,10}

	for i, elem := range arr {
		fmt.Println("Dentro")
		fmt.Println(i, elem)

	}
	for _, elem := range arr {
		fmt.Println("Dentro")
		fmt.Println(elem)

	}


	//novo
	for range 10 {
		//mesma coisa que o padrao tradicional
	}

	for i:= range 10 {
		//mesma coisa que o padrao tradicional
		fmt.Println(i)
	}

	fmt.Println(result)

	
}