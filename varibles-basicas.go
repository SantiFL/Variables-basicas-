package main

import (
	"fmt"
	"time"
)

func main() {
	var suma int = 10 + 10
	var resta int = 10 - 5
	var NombreUsuario = "Santiago"
	time.Sleep(time.Second * 10)
	fmt.Println("hola buenas tardes " + NombreUsuario)
	fmt.Println(suma)
	fmt.Println(resta)
}
