package x1

import "fmt"
import abstractproduct "design-patterns/abstract-factory/iproduct"

// ProductX1B 具体X1系列产品B
type ProductX1B struct {
	ID    string
	Name  string
	Price int32
}

// NewProductX1B 创建X1系列产品B
func NewProductX1B() abstractproduct.ProductB {
	return &ProductX1B{
		ID:    "B-X1.101",
		Name:  "ProductX1-B",
		Price: 120,
	}
}

// GetProductName 获取产品名
func (p *ProductX1B) GetProductName() string {
	return p.Name
}

// GetProductID 获取产品ID
func (p *ProductX1B) GetProductID() string {
	return p.ID
}

// GetProductPrice 获取产品价格
func (p *ProductX1B) GetProductPrice() int32 {
	return p.Price
}

// PrintProductPrice 打印产品价格
func (p *ProductX1B) PrintProductPrice() {
	fmt.Printf("The Price of Product B is %d", p.Price)
}

// PrintProductPrice 打印产品价格
func (p *ProductX1B) PrintProductPrices() {
	fmt.Printf("The Price of Product B is %d", p.Price)
}
