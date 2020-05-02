package abstractfactory

import (
	abstractproduct "design-patterns/abstract-factory/iproduct"
)

//Factory 创建产品的工厂接口
type Factory interface {
	CreateProductA() abstractproduct.ProductA
	CreateProductB() abstractproduct.ProductB
}
