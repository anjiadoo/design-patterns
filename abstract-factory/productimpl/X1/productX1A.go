package x1

import "fmt"
import abstractproduct "design-patterns/abstract-factory/iproduct"

// ProductX1A 具体X1系列产品A
type ProductX1A struct {
	ID    string
	Name  string
	Price int32
}

// NewProductX1A 创建X1系列产品A
func NewProductX1A() abstractproduct.ProductA {
	return &ProductX1A{
		ID:    "A-X1.101",
		Name:  "ProductX1-A",
		Price: 100,
	}
}

// GetProductName 获取产品名
func (p *ProductX1A) GetProductName() string {
	return p.Name
}

// GetProductID 获取产品ID
func (p *ProductX1A) GetProductID() string {
	return p.ID
}

// GetProductPrice 获取产品价格
func (p *ProductX1A) GetProductPrice() int32 {
	return p.Price
}

// PrintProductPrice 打印产品价格
func (p *ProductX1A) PrintProductPrice() {
	fmt.Printf("The Price of Product A is %d", p.Price)
}
