package entity

type Games struct {
	GameID      int
	Title       string
	Description string
	Price       float64
	Developer   string
	Publisher   string
	Rating      float64
}

type Publisher struct {
	Name     string
	TotalBuy int
}

type TopGame struct {
	Name     string
	TotalBuy int
}
