package main

import (
	"fmt"
	"github.com/lozovoya/gohomework5_1/pkg/card"
	"github.com/lozovoya/gohomework5_1/pkg/transfer"
)

func main() {

	png := *card.NewService("Penguin Bank")
	png.IssueCard("master", 100_000_00, "5106 2100 0000 0000", "rub")
	png.IssueCard("visa", 100_000_00, "5106 2111 1111 1111", "rub")
	png.IssueCard("master", 10_000_00, "5106 2122 2222 2222", "rub")
	png.IssueCard("visa", 15_000_00, "5106 2133 3333 3333", "rub")
	png.IssueCard("master", 50_000_00, "5106 2144 4444 4444", "rub")
	png.IssueCard("visa", 60_000_00, "5106 2155 5555 5555", "rub")

	pngTr := *transfer.NewService(&png,
		0, 0,
		5, 10_00,
		0, 0,
		15, 30_00)

	fmt.Println(pngTr.Card2Card("5106 2111 1111 1111", "5106 2100 0000 0000", 1_000_00))
	err := pngTr.Transfer("5106 2111 1111 1111", "5106 2100 0000 0000", 1_000_00)
	if err != nil {
		switch err {
		case transfer.ErrorSourceCardNotEnoughMoney:
			fmt.Println("Cannot complete transfer")
		case transfer.ErrorSourceCardNotFound:
			fmt.Println("Check from card")
		case transfer.ErrorDestCardNotFound:
			fmt.Println("Check destination card")
		default:
			fmt.Println("General Error")
		}

	}
}
