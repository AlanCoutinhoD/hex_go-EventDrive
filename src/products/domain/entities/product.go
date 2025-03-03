package entities

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Image       string
}

func NewProduct(name string, price float64, description string, image string) *Product {
	return &Product{
		ID:          1,
		Name:        name,
		Price:       price,
		Description: description,
		Image:       image,
	}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) SetPrice(price float64) {
	p.Price = price
}

func (p *Product) GetID() int {
	return p.ID
}

func (p *Product) SetID(id int) {
	p.ID = id
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) SetDescription(description string) {
	p.Description = description
}

func (p *Product) GetImage() string {
	return p.Image
}

func (p *Product) SetImage(image string) {
	p.Image = image
}

func (p *Product) Delete() {
	p.ID = 0
	p.Name = ""
	p.Price = 0
	p.Description = ""
	p.Image = ""
}
