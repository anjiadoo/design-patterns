package x2

import "fmt"
import abstractproduct "design-patterns/abstract-factory/iproduct"

// ProductX2A 具体X2系列产品A
type ProductX2A struct {
	ID    string
	Name  string
	Price int32
}

// NewProductX2A 创建X2系列产品A
func NewProductX2A() abstractproduct.ProductA {
	return &ProductX2A{
		ID:    "A-X2.101",
		Name:  "ProductX2-A",
		Price: 200,
	}
}

// GetProductName 获取产品名
func (p *ProductX2A) GetProductName() string {
	return p.Name
}

// GetProductID 获取产品ID
func (p *ProductX2A) GetProductID() string {
	return p.ID
}

// GetProductPrice 获取产品价格
func (p *ProductX2A) GetProductPrice() int32 {
	return p.Price
}

// PrintProductPrice 打印产品价格
func (p *ProductX2A) PrintProductPrice() {
	fmt.Printf("The Price of Product B is %d", p.Price)
}
