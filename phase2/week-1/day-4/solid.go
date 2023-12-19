package main

// Class atau modeule yang hanya memiliki sebuah tanggung jawab
// jangan gabungkan sebuah fungsi dalam suatu entitas
// lebih dr 2 method bergantung pd sebuah struct
import "fmt"

type Order struct {
	Item []string
}

func main() {
	fmt.Println("Hii")
}
