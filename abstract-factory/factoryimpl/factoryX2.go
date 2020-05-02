package concretefactory

import (
	abstractfactory "design-patterns/abstract-factory/ifactory"
	abstractproduct "design-patterns/abstract-factory/iproduct"
	x2 "design-patterns/abstract-factory/productimpl/X2"
)

//FactoryX2 X2系列的工厂
type FactoryX2 struct{}

//GetFactoryX2 创建X1系列工厂
func GetFactoryX2() abstractfactory.Factory {
	return new(FactoryX2)
}

//CreateProductA 创建X2系列产品A
func (f *FactoryX2) CreateProductA() abstractproduct.ProductA {
	return x2.NewProductX2A()
}

//CreateProductB 创建X2系列产品B
func (f *FactoryX2) CreateProductB() abstractproduct.ProductB {
	return x2.NewProductX2B()
}
