package entity

type Product struct {
	ProductId   int
	ProductName string
	Category    string
	Price       float64
}

type Addcart struct {
	CartId    int
	UserId    int
	ProductId int
	Quantity  int
}

type ListCart struct {
	UserId      int
	ProductId   int
	ProductName string
	Quantity    int
}