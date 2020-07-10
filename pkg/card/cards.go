package card

import (
	"errors"
	"strconv"
	"strings"
)

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

var ErrorWrongCardNumber = errors.New("Check card number")

func (s *Service) IssueCard(issuer string, balance int64, number string, currency string) *Card {
	card := &Card{
		Id:       int64(len(s.Cards)),
		Issuer:   issuer,
		Balance:  balance,
		Currency: currency,
		Number:   number,
		Icon:     "http://....",
	}
	s.Cards = append(s.Cards, card)
	return card
}

func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}

func (s *Service) FindByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}

func IsValid(cardNumber string) error {
	cardstr := strings.Split(strings.ReplaceAll(cardNumber, " ", ""), "")
	cardint := make([]int, len(cardstr), len(cardstr))
	for i, j := range cardstr {
		cardint[i], _ = strconv.Atoi(j)
	}

	for i := 1; i < len(cardint); i = i + 2 {
		if cardint[i]*2 > 9 {
			cardint[i] = cardint[i]*2 - 9
		} else {
			cardint[i] = cardint[i] * 2
		}
	}
	sum := 0
	for _, i := range cardint {
		sum += i
	}

	if (sum % 10) == 0 {
		return nil
	}

	return ErrorWrongCardNumber
}
