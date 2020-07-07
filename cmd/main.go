package main

import (
	"fmt"
	"github.com/lozovoya/gohomework5_1/pkg/card"
	"github.com/lozovoya/gohomework5_1/pkg/transfer"
)

func main() {

	png := *card.NewService("Penguin Bank")
	png.IssueCard("master", 100_000_00, "0000 0000 0000 0000", "rub")
	png.IssueCard("visa", 100_000_00, "1111 1111 1111 1111", "rub")
	png.IssueCard("master", 10_000_00, "2222 2222 2222 2222", "rub")
	png.IssueCard("visa", 15_000_00, "3333 3333 3333 3333", "rub")
	png.IssueCard("master", 50_000_00, "4444 4444 4444 4444", "rub")
	png.IssueCard("visa", 60_000_00, "5555 5555 5555 5555", "rub")

	pngTr := *transfer.NewService(&png,
		0, 0,
		5, 10_00,
		0, 0,
		15, 30_00)

	fmt.Println(pngTr.Card2Card("1111 1111 1111 1111", "000 0000 0000 0000", 1_000_00))
}
