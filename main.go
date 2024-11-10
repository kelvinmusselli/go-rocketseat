package main

import (
	"fmt"
	"myFirstGoProject/pacote"
)

func main(){
	fmt.Println("Hello World")
	fmt.Println(pacote.Foo)
	pacote.PrintMinha()
}

//se a primeira letra da funcao ou var for
//minuscula é privada se for mais maiuscula
// é publica