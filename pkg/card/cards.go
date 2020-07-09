package card

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

//func (s *Service) FindByNumber(number string) (*Card, bool) {
//	if strings.HasPrefix(number, "510621") {
//		for _, card := range s.Cards {
//			if card.Number == number {
//				return card, true
//			}
//		}
//	}
//	return nil, false
//}

func (s *Service) FindByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}
