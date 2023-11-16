package main

import "fmt"

func main() {
	idealWeight(50, 180)
}

func idealWeight(weight, height int) {
	idWe := height - 110
	if weight == idWe {
		fmt.Println("Ты мощны")
	} else if weight > idWe {
		fmt.Println("Жирдяй")
		fmt.Println("Сбрось ", weight-idWe, "кг")
	} else if weight < idWe {
		fmt.Println("Дрищ")
		fmt.Println("Набери", idWe-weight, "кг")
	}
}
