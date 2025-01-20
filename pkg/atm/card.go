package atm

type Card struct {
	cardNo int
	pin    int
}

func NewCard(cardNo int, pin int) *Card {
	return &Card{
		cardNo: cardNo,
		pin:    pin,
	}
}
