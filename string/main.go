package main

import (
	"fmt"
	"strings"
)

func main() {
	product := "Kayak"
	//checks sub string
	fmt.Println("contains :", strings.Contains(product, "yak"))
	fmt.Println("contains any :", strings.ContainsAny(product, "ab"))
	fmt.Println("conytains Rune", strings.ContainsRune(product, 'a'))
	fmt.Println("Equal fold :", strings.EqualFold(product, "KAYAK"))
	fmt.Println("Has prefix :", strings.HasPrefix(product, "Ka"))
	fmt.Println("Has suffix :", strings.HasSuffix(product, "yak"))
}
