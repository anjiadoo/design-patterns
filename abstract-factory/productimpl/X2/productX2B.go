package x2

import "fmt"
import abstractproduct "design-patterns/abstract-factory/iproduct"

// ProductX2B 具体X2系列产品B
type ProductX2B struct {
	ID    string
	Name  string
	Price int32
}

// NewProductX2B 创建X2系列产品B2
func NewProductX2B() abstractproduct.ProductB {
	return &ProductX2B{
		ID:    "B-X2.101",
		Name:  "ProductX2-B",
		Price: 240,
	}
}

// GetProductName 获取产品名
func (p *ProductX2B) GetProductName() string {
	return p.Name
}

// GetProductID 获取产品ID
func (p *ProductX2B) GetProductID() string {
	return p.ID
}

// GetProductPrice 获取产品价格
func (p *ProductX2B) GetProductPrice() int32 {
	return p.Price
}

// PrintProductPrice 打印产品价格
func (p *ProductX2B) PrintProductPrice() {
	fmt.Printf("The Price of Product B is %d", p.Price)
}

// PrintProductPrice 打印产品价格
func (p *ProductX2B) PrintProductPrices() {
	fmt.Printf("The Price of Product B is %d", p.Price)
}
