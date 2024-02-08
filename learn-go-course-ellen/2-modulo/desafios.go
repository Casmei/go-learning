package main

import (
	"fmt"
)

const irmao string = "Maria"
const irma = "João"

const (
	_ = 2024 + iota
	ANO_1
	ANO_2
	ANO_3
	ANO_4
)

func main() {
	fmt.Println("--------------")
	exercicioUm()
	fmt.Println("--------------")
	exercicioDois()
	fmt.Println("--------------")
	exercicioTres()
	fmt.Println("--------------")
	exercicioQuatro()
	fmt.Println("--------------")
	exercicioCinco()
	fmt.Println("--------------")
	exercicioSeis()
	fmt.Println("--------------")
}

func exercicioUm() {
	valor := 42

	fmt.Printf("decimal:%d\tbinario:%b\thexadecimal:%#x\n", valor, valor, valor)
}

func exercicioDois() {
	d := 53 == 53
	a := 42 != 41
	b := 1 <= -4
	f := 5 < 3
	e := 5 >= 10
	c := 3 > 3

	fmt.Printf("a:%v \tb:%v \tc:%v \td:%v \te:%v \tf:%v\n", a, b, c, d, e, f)
}

func exercicioTres() {
	fmt.Printf("%v & %v\n", irmao, irma)
}

func exercicioQuatro() {
	x := 4
	y := x >> 1
	fmt.Printf("decimal:%d\tbinario:%b\thexadecimal:%#x\n", x, x, x)
	fmt.Printf("decimal:%d\tbinario:%b\thexadecimal:%#x\n", y, y, y)
}

func exercicioCinco() {
	x := `Olá,
	
	esse		rolÊ de 	
		raw						string é 		muito d
					oido!`
	fmt.Println(x)
}

func exercicioSeis() {
	// Tive certa dificuldade nesse exercisio
	fmt.Println(ANO_1, ANO_2, ANO_3, ANO_4)
}
