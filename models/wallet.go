package models

type Currency string

const (
	NgnCurrency Currency = "NGN"

	UsdCurrency Currency = "USD"
)

type Wallet struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	Currency Currency  `json:"currency"`
}

type Transaction struct {
	WalletID uuid.UUID `json:"wallet_id"`
	Debit    float64   `json:"debit"`
	Credit   float64   `json:"credit"`
}

type (
	Naira struct {
		MainUnit
	}
	Kobo struct {
		SubUnit
	}
	MainUnit float64
	SubUnit  float64
)

func (ngn Naira) ToSubUnit() float64 {
	return float64(ngn.MainUnit * 100)
}

func (k Kobo) ToMainUnit() float64 {
	return float64(k.SubUnit / 100)
}
