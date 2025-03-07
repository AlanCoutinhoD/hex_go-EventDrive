package entities

type Order struct {
	ID        int
	IdProduct int
	IdClient  string
	Quantity  int
}

func NewOrder(idProduct int, idClient string, quantity int) *Order {
	return &Order{
		ID:        0,
		IdProduct: idProduct,
		IdClient:  idClient,
		Quantity:  quantity,
	}
}

func (o *Order) GetID() int {
	return o.ID
}

func (o *Order) SetID(id int) {
	o.ID = id
}

func (o *Order) GetIdProduct() int {
	return o.IdProduct
}

func (o *Order) SetIdProduct(idProduct int) {
	o.IdProduct = idProduct
}

func (o *Order) GetIdClient() string {
	return o.IdClient
}

func (o *Order) SetIdClient(idClient string) {
	o.IdClient = idClient
}

func (o *Order) GetQuantity() int {
	return o.Quantity
}

func (o *Order) SetQuantity(quantity int) {
	o.Quantity = quantity
}
