package concretefactory

import (
	abstractfactory "design-patterns/abstract-factory/ifactory"
	abstractproduct "design-patterns/abstract-factory/iproduct"
	x1 "design-patterns/abstract-factory/productimpl/X1"
)

//FactoryX1 X1系列的工厂
type FactoryX1 struct{}

//GetFactoryX1 创建X1系列工厂
func GetFactoryX1() abstractfactory.Factory {
	return new(FactoryX1)
}

//CreateProductA 创建X1系列产品A
func (f *FactoryX1) CreateProductA() abstractproduct.ProductA {
	return x1.NewProductX1A()
}

//CreateProductB 创建X1系列产品B
func (f *FactoryX1) CreateProductB() abstractproduct.ProductB {
	return x1.NewProductX1B()
}
