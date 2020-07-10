package transfer

import (
	"errors"
	"github.com/lozovoya/gohomework5_1/pkg/card"
	"strings"
)

type Service struct {
	CardSvc       *card.Service
	ItoICommision int64
	ItoIMin       int64
	ItoECommision int64
	ItoEMin       int64
	EtoICommision int64
	EtoIMin       int64
	EtoECommision int64
	EtoEMin       int64
}

var ErrorSourceCardNotEnoughMoney = errors.New("source card: not enough money")
var ErrorSourceCardNotFound = errors.New("source card: not found")
var ErrorDestCardNotFound = errors.New("destination card: not found")

const bankCode = "5106 21"

func NewService(cardSvc *card.Service, itoICommision int64, itoIMin int64, itoECommision int64, itoEMin int64,
	etoICommision int64, etoIMin int64, etoECommision int64, etoEMin int64) *Service {
	return &Service{CardSvc: cardSvc,
		ItoICommision: itoICommision, ItoIMin: itoIMin,
		ItoECommision: itoECommision, ItoEMin: itoEMin,
		EtoICommision: etoICommision, EtoIMin: etoIMin,
		EtoECommision: etoECommision, EtoEMin: etoEMin}
}

func (s *Service) Card2Card(from string, to string, amount int64) (int64, error) {
	fromCard := s.CardSvc.FindByNumber(from)
	toCard := s.CardSvc.FindByNumber(to)

	if strings.HasPrefix(from, bankCode) && (fromCard == nil) {
		return 0, ErrorSourceCardNotFound
	}
	if strings.HasPrefix(to, bankCode) && (toCard == nil) {
		return 0, ErrorDestCardNotFound
	}

	// I to I
	if (fromCard != nil) && (toCard != nil) {
		if fromCard.Balance < amount {
			return 0, ErrorSourceCardNotEnoughMoney
		}
		fromCard.Balance -= amount
		toCard.Balance += amount
		return amount, nil
	}

	// I to E
	if (fromCard != nil) && (toCard == nil) {
		commission := amount * s.ItoECommision / 1000
		total := amount + commission
		if fromCard.Balance < total {
			return 0, ErrorSourceCardNotEnoughMoney
		}
		fromCard.Balance -= total
		return total, nil
	}

	// E to I
	if (fromCard == nil) && (toCard != nil) {
		toCard.Balance += amount
		return 0, nil
	}

	// E to E
	if (fromCard == nil) && (toCard == nil) {
		commission := amount * s.EtoECommision / 1000
		if commission > s.EtoEMin {
			total := amount + commission
			return total, nil
		}
		total := amount + s.EtoEMin
		return total, nil
	}
	return 0, nil
}
