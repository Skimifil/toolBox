package main

import (
	"fmt"
	"strconv"
)

type Menu struct {
	opcao int
	path  string
}

func CreateMenu() Menu {
	menu := Menu{
		opcao: 1,
		path:  "opcao/menus/GetListEventsActive",
	}
	return menu
}

func ShowMenu(opcao int) func(string) {
	return func(path string) {
		if opcao == 1 {
			fmt.Println("Opção: " + strconv.Itoa(opcao) + " Escolhido, com isso foi chamado o path " + path)
		} else if opcao == 2 {
			fmt.Println("Opção: " + strconv.Itoa(opcao) + " Escolhido, com isso foi chamado o path " + path)
		}
	}
}

func main() {
	fmt.Println("Hello World")
	x := CreateMenu()
	fmt.Println(x.opcao)
	fmt.Println(x.path)

	y := ShowMenu(x.opcao)
	y(x.path)
	y("2")
}
