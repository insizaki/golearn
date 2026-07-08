package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	Number string
	Type   string
}

func main() {
	var listCard []Card
	var input int
	fmt.Scan(&input)
	str := strconv.Itoa(input)
	s := str[:2]

	if s == "62" && len(str) >= 16 && len(str) <= 19 {
		listCard = append(listCard, Card{Number: str, Type: "China UnionPay"})
	} else if str[:4] == "4903" || str[:4] == "4905" || str[:4] == "4911" || str[:4] == "4936" || str[:6] == "564182" || str[:6] == "63310" || str[:4] == "6333" || len(str) == 18 || len(str) == 16 || len(str) == 19 {
		listCard = append(listCard, Card{Number: str, Type: "Switch"})
	}

	fmt.Println("Nomor ", listCard[0].Number, " adalah ", listCard[0].Type)
}
