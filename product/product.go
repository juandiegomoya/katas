package product

type Product struct {
	name     string
	price    float64
	quantity int
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetQuantity(quantity int) {
	p.quantity = quantity
}

func (p *Product) GetQuantity() int {
	return p.quantity
}
